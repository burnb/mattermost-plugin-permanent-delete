// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testFileinfos(t *testing.T) {
	t.Parallel()

	query := Fileinfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFileinfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFileinfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Fileinfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFileinfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FileinfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFileinfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FileinfoExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Fileinfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FileinfoExists to return true, but got false.")
	}
}

func testFileinfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	fileinfoFound, err := FindFileinfo(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if fileinfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFileinfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Fileinfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFileinfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Fileinfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFileinfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	fileinfoOne := &Fileinfo{}
	fileinfoTwo := &Fileinfo{}
	if err = randomize.Struct(seed, fileinfoOne, fileinfoDBTypes, false, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}
	if err = randomize.Struct(seed, fileinfoTwo, fileinfoDBTypes, false, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fileinfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fileinfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Fileinfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFileinfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	fileinfoOne := &Fileinfo{}
	fileinfoTwo := &Fileinfo{}
	if err = randomize.Struct(seed, fileinfoOne, fileinfoDBTypes, false, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}
	if err = randomize.Struct(seed, fileinfoTwo, fileinfoDBTypes, false, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fileinfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fileinfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func fileinfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func fileinfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func fileinfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func fileinfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func fileinfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func fileinfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func fileinfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func fileinfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func fileinfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Fileinfo) error {
	*o = Fileinfo{}
	return nil
}

func testFileinfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Fileinfo{}
	o := &Fileinfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, fileinfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Fileinfo object: %s", err)
	}

	AddFileinfoHook(boil.BeforeInsertHook, fileinfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	fileinfoBeforeInsertHooks = []FileinfoHook{}

	AddFileinfoHook(boil.AfterInsertHook, fileinfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	fileinfoAfterInsertHooks = []FileinfoHook{}

	AddFileinfoHook(boil.AfterSelectHook, fileinfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	fileinfoAfterSelectHooks = []FileinfoHook{}

	AddFileinfoHook(boil.BeforeUpdateHook, fileinfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	fileinfoBeforeUpdateHooks = []FileinfoHook{}

	AddFileinfoHook(boil.AfterUpdateHook, fileinfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	fileinfoAfterUpdateHooks = []FileinfoHook{}

	AddFileinfoHook(boil.BeforeDeleteHook, fileinfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	fileinfoBeforeDeleteHooks = []FileinfoHook{}

	AddFileinfoHook(boil.AfterDeleteHook, fileinfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	fileinfoAfterDeleteHooks = []FileinfoHook{}

	AddFileinfoHook(boil.BeforeUpsertHook, fileinfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	fileinfoBeforeUpsertHooks = []FileinfoHook{}

	AddFileinfoHook(boil.AfterUpsertHook, fileinfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	fileinfoAfterUpsertHooks = []FileinfoHook{}
}

func testFileinfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFileinfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(fileinfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFileinfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFileinfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FileinfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFileinfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Fileinfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	fileinfoDBTypes = map[string]string{`ID`: `character varying`, `Creatorid`: `character varying`, `Postid`: `character varying`, `Createat`: `bigint`, `Updateat`: `bigint`, `Deleteat`: `bigint`, `Path`: `character varying`, `Thumbnailpath`: `character varying`, `Previewpath`: `character varying`, `Name`: `character varying`, `Extension`: `character varying`, `Size`: `bigint`, `Mimetype`: `character varying`, `Width`: `integer`, `Height`: `integer`, `Haspreviewimage`: `boolean`}
	_               = bytes.MinRead
)

func testFileinfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(fileinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(fileinfoAllColumns) == len(fileinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFileinfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(fileinfoAllColumns) == len(fileinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Fileinfo{}
	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fileinfoDBTypes, true, fileinfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(fileinfoAllColumns, fileinfoPrimaryKeyColumns) {
		fields = fileinfoAllColumns
	} else {
		fields = strmangle.SetComplement(
			fileinfoAllColumns,
			fileinfoPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := FileinfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFileinfosUpsert(t *testing.T) {
	t.Parallel()

	if len(fileinfoAllColumns) == len(fileinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Fileinfo{}
	if err = randomize.Struct(seed, &o, fileinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Fileinfo: %s", err)
	}

	count, err := Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, fileinfoDBTypes, false, fileinfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Fileinfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Fileinfo: %s", err)
	}

	count, err = Fileinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}