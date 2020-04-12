package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/mattermost/mattermost-plugin-starter-template/models"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"os"
	"sync"
	"time"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	db *sql.DB
}

func (p *Plugin) OnActivate() error {
	currentDir, _ := os.Getwd()
	p.API.LogDebug(fmt.Sprintf("Current Path:%s", currentDir))
	var err error
	if err := envconfig.Process("MM", p.configuration); err != nil {
		return fmt.Errorf("CAN'T LOAD CONFIG: %v", err)
	}
	p.API.LogDebug(fmt.Sprintf("DatatSource:%s", p.configuration.DataSource))
	p.API.LogDebug(fmt.Sprintf("FileStorageDirectory:%s", p.configuration.FileStorageDirectory))
	if _, err := os.Stat(p.configuration.FileStorageDirectory); err != nil && os.IsNotExist(err) {
		return fmt.Errorf("FILE STORAGE DOESN'T EXIST('%s')", p.configuration.FileStorageDirectory)
	}
	if p.db, err = sql.Open("postgres", p.configuration.DataSource); err != nil {
		return fmt.Errorf("CAN'T OPEN DB: %v", err)
	}
	ctx := context.Background()
	if err = p.permanentPostsDelete(ctx); err != nil {
		return err
	}
	go func() {
		for {
			<-time.After(time.Duration(p.getConfiguration().Waiting) * time.Second)
			_ = p.permanentPostsDelete(ctx)
		}
	}()

	return nil
}

func (p *Plugin) permanentPostDelete(ctx context.Context, post *models.Post) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	files, err := models.Fileinfos(qm.Where("postid = ?", post.ID)).All(ctx, p.db)
	if err != nil {
		return err
	}
	for _, file := range files {
		p.API.LogDebug(fmt.Sprintf("Deleting file #%s: %s", file.ID, p.fullFilePath(file.Path.String)))
		if err = os.Remove(p.fullFilePath(file.Path.String)); err != nil {
			p.API.LogError(fmt.Sprintf("CAN'T REMOVE FILE #%s: %v", file.ID, err))
		}
		if file.Thumbnailpath.String != "" {
			if err = os.Remove(p.fullFilePath(file.Thumbnailpath.String)); err != nil {
				p.API.LogError(fmt.Sprintf("СAN'T REMOVE FILE THUMBNAIL #%s: %v", file.ID, err))
			}
		}
		if file.Previewpath.String != "" {
			if err = os.Remove(p.fullFilePath(file.Previewpath.String)); err != nil {
				p.API.LogError(fmt.Sprintf("СAN'T REMOVE FILE PREVIEW #%s: %v", file.ID, err))
			}
		}
		if _, err = file.Delete(ctx, p.db); err != nil {
			_ = tx.Rollback()

			return fmt.Errorf("сan't delete fileinfo #%s: %v", file.ID, err)
		}
	}
	if _, err = models.Reactions(models.ReactionWhere.Postid.EQ(post.ID)).DeleteAll(ctx, p.db); err != nil {
		_ = tx.Rollback()

		return fmt.Errorf("сan't delete post reactions: %v", err)
	}
	_, err =
		models.Preferences(
			qm.Expr(
				models.PreferenceWhere.Category.EQ("flagged_post"),
				models.PreferenceWhere.Name.EQ(post.ID),
			),
		).DeleteAll(ctx, p.db)
	if err != nil {
		_ = tx.Rollback()

		return fmt.Errorf("сan't delete flagged reactions: %v", err)
	}
	if _, err := post.Delete(ctx, p.db); err != nil {
		_ = tx.Rollback()

		return err
	}
	_ = tx.Commit()
	p.API.LogDebug(fmt.Sprintf("Post #%s deleted", post.ID))

	return nil
}

func (p *Plugin) permanentPostsDelete(ctx context.Context) error {
	p.API.LogDebug("Running...")
	posts, err := models.Posts(qm.Where("deleteat != ?", 0)).All(ctx, p.db)
	if err != nil {
		return err
	}
	for _, post := range posts {
		if err := p.permanentPostDelete(ctx, post); err != nil {
			p.API.LogError(fmt.Sprintf("CAN'T PERMANENT DELETE POST #%s: %v", post.ID, err))
		}
	}
	p.API.LogDebug("Waiting...")

	return nil
}

func (p *Plugin) fullFilePath(filePath string) string {
	return p.getConfiguration().FileStorageDirectory + filePath
}
