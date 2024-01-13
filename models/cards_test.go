// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testCards(t *testing.T) {
	t.Parallel()

	query := Cards()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCardsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
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

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCardsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Cards().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCardsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CardSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCardsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CardExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Card exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CardExists to return true, but got false.")
	}
}

func testCardsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	cardFound, err := FindCard(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if cardFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCardsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Cards().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCardsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Cards().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCardsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cardOne := &Card{}
	cardTwo := &Card{}
	if err = randomize.Struct(seed, cardOne, cardDBTypes, false, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}
	if err = randomize.Struct(seed, cardTwo, cardDBTypes, false, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Cards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCardsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cardOne := &Card{}
	cardTwo := &Card{}
	if err = randomize.Struct(seed, cardOne, cardDBTypes, false, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}
	if err = randomize.Struct(seed, cardTwo, cardDBTypes, false, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func cardBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func cardAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func cardAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func cardBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func cardAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func cardBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func cardAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func cardBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func cardAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Card) error {
	*o = Card{}
	return nil
}

func testCardsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Card{}
	o := &Card{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cardDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Card object: %s", err)
	}

	AddCardHook(boil.BeforeInsertHook, cardBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cardBeforeInsertHooks = []CardHook{}

	AddCardHook(boil.AfterInsertHook, cardAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cardAfterInsertHooks = []CardHook{}

	AddCardHook(boil.AfterSelectHook, cardAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cardAfterSelectHooks = []CardHook{}

	AddCardHook(boil.BeforeUpdateHook, cardBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cardBeforeUpdateHooks = []CardHook{}

	AddCardHook(boil.AfterUpdateHook, cardAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cardAfterUpdateHooks = []CardHook{}

	AddCardHook(boil.BeforeDeleteHook, cardBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cardBeforeDeleteHooks = []CardHook{}

	AddCardHook(boil.AfterDeleteHook, cardAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cardAfterDeleteHooks = []CardHook{}

	AddCardHook(boil.BeforeUpsertHook, cardBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cardBeforeUpsertHooks = []CardHook{}

	AddCardHook(boil.AfterUpsertHook, cardAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cardAfterUpsertHooks = []CardHook{}
}

func testCardsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCardsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(cardColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCardToManyCard1Battles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Card
	var b, c Battle

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, battleDBTypes, false, battleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, battleDBTypes, false, battleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.Card1ID = a.ID
	c.Card1ID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Card1Battles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.Card1ID == b.Card1ID {
			bFound = true
		}
		if v.Card1ID == c.Card1ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CardSlice{&a}
	if err = a.L.LoadCard1Battles(ctx, tx, false, (*[]*Card)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Card1Battles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Card1Battles = nil
	if err = a.L.LoadCard1Battles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Card1Battles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCardToManyCard2Battles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Card
	var b, c Battle

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, battleDBTypes, false, battleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, battleDBTypes, false, battleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.Card2ID = a.ID
	c.Card2ID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Card2Battles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.Card2ID == b.Card2ID {
			bFound = true
		}
		if v.Card2ID == c.Card2ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CardSlice{&a}
	if err = a.L.LoadCard2Battles(ctx, tx, false, (*[]*Card)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Card2Battles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Card2Battles = nil
	if err = a.L.LoadCard2Battles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Card2Battles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCardToManyAddOpCard1Battles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Card
	var b, c, d, e Battle

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cardDBTypes, false, strmangle.SetComplement(cardPrimaryKeyColumns, cardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Battle{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, battleDBTypes, false, strmangle.SetComplement(battlePrimaryKeyColumns, battleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Battle{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCard1Battles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.Card1ID {
			t.Error("foreign key was wrong value", a.ID, first.Card1ID)
		}
		if a.ID != second.Card1ID {
			t.Error("foreign key was wrong value", a.ID, second.Card1ID)
		}

		if first.R.Card1 != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Card1 != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Card1Battles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Card1Battles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Card1Battles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCardToManyAddOpCard2Battles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Card
	var b, c, d, e Battle

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cardDBTypes, false, strmangle.SetComplement(cardPrimaryKeyColumns, cardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Battle{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, battleDBTypes, false, strmangle.SetComplement(battlePrimaryKeyColumns, battleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Battle{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCard2Battles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.Card2ID {
			t.Error("foreign key was wrong value", a.ID, first.Card2ID)
		}
		if a.ID != second.Card2ID {
			t.Error("foreign key was wrong value", a.ID, second.Card2ID)
		}

		if first.R.Card2 != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Card2 != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Card2Battles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Card2Battles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Card2Battles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCardsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
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

func testCardsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CardSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCardsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Cards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cardDBTypes = map[string]string{`ID`: `integer`, `Wins`: `integer`, `Battles`: `integer`, `Name`: `text`, `Token`: `text`, `Filename`: `text`, `Accepted`: `boolean`}
	_           = bytes.MinRead
)

func testCardsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(cardPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(cardAllColumns) == len(cardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cardDBTypes, true, cardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCardsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cardAllColumns) == len(cardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Card{}
	if err = randomize.Struct(seed, o, cardDBTypes, true, cardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cardDBTypes, true, cardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cardAllColumns, cardPrimaryKeyColumns) {
		fields = cardAllColumns
	} else {
		fields = strmangle.SetComplement(
			cardAllColumns,
			cardPrimaryKeyColumns,
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

	slice := CardSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCardsUpsert(t *testing.T) {
	t.Parallel()

	if len(cardAllColumns) == len(cardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Card{}
	if err = randomize.Struct(seed, &o, cardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Card: %s", err)
	}

	count, err := Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, cardDBTypes, false, cardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Card struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Card: %s", err)
	}

	count, err = Cards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}