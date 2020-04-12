// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Reaction is an object representing the database table.
type Reaction struct {
	Userid    string     `boil:"userid" json:"userid" toml:"userid" yaml:"userid"`
	Postid    string     `boil:"postid" json:"postid" toml:"postid" yaml:"postid"`
	Emojiname string     `boil:"emojiname" json:"emojiname" toml:"emojiname" yaml:"emojiname"`
	Createat  null.Int64 `boil:"createat" json:"createat,omitempty" toml:"createat" yaml:"createat,omitempty"`

	R *reactionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L reactionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReactionColumns = struct {
	Userid    string
	Postid    string
	Emojiname string
	Createat  string
}{
	Userid:    "userid",
	Postid:    "postid",
	Emojiname: "emojiname",
	Createat:  "createat",
}

// Generated where

var ReactionWhere = struct {
	Userid    whereHelperstring
	Postid    whereHelperstring
	Emojiname whereHelperstring
	Createat  whereHelpernull_Int64
}{
	Userid:    whereHelperstring{field: "\"reactions\".\"userid\""},
	Postid:    whereHelperstring{field: "\"reactions\".\"postid\""},
	Emojiname: whereHelperstring{field: "\"reactions\".\"emojiname\""},
	Createat:  whereHelpernull_Int64{field: "\"reactions\".\"createat\""},
}

// ReactionRels is where relationship names are stored.
var ReactionRels = struct {
}{}

// reactionR is where relationships are stored.
type reactionR struct {
}

// NewStruct creates a new relationship struct
func (*reactionR) NewStruct() *reactionR {
	return &reactionR{}
}

// reactionL is where Load methods for each relationship are stored.
type reactionL struct{}

var (
	reactionAllColumns            = []string{"userid", "postid", "emojiname", "createat"}
	reactionColumnsWithoutDefault = []string{"userid", "postid", "emojiname", "createat"}
	reactionColumnsWithDefault    = []string{}
	reactionPrimaryKeyColumns     = []string{"userid", "postid", "emojiname"}
)

type (
	// ReactionSlice is an alias for a slice of pointers to Reaction.
	// This should generally be used opposed to []Reaction.
	ReactionSlice []*Reaction
	// ReactionHook is the signature for custom Reaction hook methods
	ReactionHook func(context.Context, boil.ContextExecutor, *Reaction) error

	reactionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	reactionType                 = reflect.TypeOf(&Reaction{})
	reactionMapping              = queries.MakeStructMapping(reactionType)
	reactionPrimaryKeyMapping, _ = queries.BindMapping(reactionType, reactionMapping, reactionPrimaryKeyColumns)
	reactionInsertCacheMut       sync.RWMutex
	reactionInsertCache          = make(map[string]insertCache)
	reactionUpdateCacheMut       sync.RWMutex
	reactionUpdateCache          = make(map[string]updateCache)
	reactionUpsertCacheMut       sync.RWMutex
	reactionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var reactionBeforeInsertHooks []ReactionHook
var reactionBeforeUpdateHooks []ReactionHook
var reactionBeforeDeleteHooks []ReactionHook
var reactionBeforeUpsertHooks []ReactionHook

var reactionAfterInsertHooks []ReactionHook
var reactionAfterSelectHooks []ReactionHook
var reactionAfterUpdateHooks []ReactionHook
var reactionAfterDeleteHooks []ReactionHook
var reactionAfterUpsertHooks []ReactionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Reaction) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Reaction) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Reaction) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Reaction) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Reaction) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Reaction) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Reaction) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Reaction) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Reaction) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reactionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddReactionHook registers your hook function for all future operations.
func AddReactionHook(hookPoint boil.HookPoint, reactionHook ReactionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		reactionBeforeInsertHooks = append(reactionBeforeInsertHooks, reactionHook)
	case boil.BeforeUpdateHook:
		reactionBeforeUpdateHooks = append(reactionBeforeUpdateHooks, reactionHook)
	case boil.BeforeDeleteHook:
		reactionBeforeDeleteHooks = append(reactionBeforeDeleteHooks, reactionHook)
	case boil.BeforeUpsertHook:
		reactionBeforeUpsertHooks = append(reactionBeforeUpsertHooks, reactionHook)
	case boil.AfterInsertHook:
		reactionAfterInsertHooks = append(reactionAfterInsertHooks, reactionHook)
	case boil.AfterSelectHook:
		reactionAfterSelectHooks = append(reactionAfterSelectHooks, reactionHook)
	case boil.AfterUpdateHook:
		reactionAfterUpdateHooks = append(reactionAfterUpdateHooks, reactionHook)
	case boil.AfterDeleteHook:
		reactionAfterDeleteHooks = append(reactionAfterDeleteHooks, reactionHook)
	case boil.AfterUpsertHook:
		reactionAfterUpsertHooks = append(reactionAfterUpsertHooks, reactionHook)
	}
}

// One returns a single reaction record from the query.
func (q reactionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Reaction, error) {
	o := &Reaction{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for reactions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Reaction records from the query.
func (q reactionQuery) All(ctx context.Context, exec boil.ContextExecutor) (ReactionSlice, error) {
	var o []*Reaction

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Reaction slice")
	}

	if len(reactionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Reaction records in the query.
func (q reactionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count reactions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q reactionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if reactions exists")
	}

	return count > 0, nil
}

// Reactions retrieves all the records using an executor.
func Reactions(mods ...qm.QueryMod) reactionQuery {
	mods = append(mods, qm.From("\"reactions\""))
	return reactionQuery{NewQuery(mods...)}
}

// FindReaction retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReaction(ctx context.Context, exec boil.ContextExecutor, userid string, postid string, emojiname string, selectCols ...string) (*Reaction, error) {
	reactionObj := &Reaction{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"reactions\" where \"userid\"=$1 AND \"postid\"=$2 AND \"emojiname\"=$3", sel,
	)

	q := queries.Raw(query, userid, postid, emojiname)

	err := q.Bind(ctx, exec, reactionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from reactions")
	}

	return reactionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Reaction) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no reactions provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(reactionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	reactionInsertCacheMut.RLock()
	cache, cached := reactionInsertCache[key]
	reactionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			reactionAllColumns,
			reactionColumnsWithDefault,
			reactionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(reactionType, reactionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(reactionType, reactionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"reactions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"reactions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into reactions")
	}

	if !cached {
		reactionInsertCacheMut.Lock()
		reactionInsertCache[key] = cache
		reactionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Reaction.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Reaction) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	reactionUpdateCacheMut.RLock()
	cache, cached := reactionUpdateCache[key]
	reactionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			reactionAllColumns,
			reactionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update reactions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"reactions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, reactionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(reactionType, reactionMapping, append(wl, reactionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update reactions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for reactions")
	}

	if !cached {
		reactionUpdateCacheMut.Lock()
		reactionUpdateCache[key] = cache
		reactionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q reactionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for reactions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for reactions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReactionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"reactions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, reactionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in reaction slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all reaction")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Reaction) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no reactions provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(reactionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	reactionUpsertCacheMut.RLock()
	cache, cached := reactionUpsertCache[key]
	reactionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			reactionAllColumns,
			reactionColumnsWithDefault,
			reactionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			reactionAllColumns,
			reactionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert reactions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(reactionPrimaryKeyColumns))
			copy(conflict, reactionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"reactions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(reactionType, reactionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(reactionType, reactionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert reactions")
	}

	if !cached {
		reactionUpsertCacheMut.Lock()
		reactionUpsertCache[key] = cache
		reactionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Reaction record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Reaction) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Reaction provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), reactionPrimaryKeyMapping)
	sql := "DELETE FROM \"reactions\" WHERE \"userid\"=$1 AND \"postid\"=$2 AND \"emojiname\"=$3"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from reactions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for reactions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q reactionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no reactionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from reactions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for reactions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReactionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(reactionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"reactions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, reactionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from reaction slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for reactions")
	}

	if len(reactionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Reaction) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindReaction(ctx, exec, o.Userid, o.Postid, o.Emojiname)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReactionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ReactionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"reactions\".* FROM \"reactions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, reactionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ReactionSlice")
	}

	*o = slice

	return nil
}

// ReactionExists checks if the Reaction row exists.
func ReactionExists(ctx context.Context, exec boil.ContextExecutor, userid string, postid string, emojiname string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"reactions\" where \"userid\"=$1 AND \"postid\"=$2 AND \"emojiname\"=$3 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userid, postid, emojiname)
	}
	row := exec.QueryRowContext(ctx, sql, userid, postid, emojiname)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if reactions exists")
	}

	return exists, nil
}