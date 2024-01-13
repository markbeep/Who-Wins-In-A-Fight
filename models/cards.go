// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Card is an object representing the database table.
type Card struct {
	ID       int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Wins     int    `boil:"wins" json:"wins" toml:"wins" yaml:"wins"`
	Battles  int    `boil:"battles" json:"battles" toml:"battles" yaml:"battles"`
	Name     string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Token    string `boil:"token" json:"token" toml:"token" yaml:"token"`
	Filename string `boil:"filename" json:"filename" toml:"filename" yaml:"filename"`
	Accepted bool   `boil:"accepted" json:"accepted" toml:"accepted" yaml:"accepted"`

	R *cardR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cardL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CardColumns = struct {
	ID       string
	Wins     string
	Battles  string
	Name     string
	Token    string
	Filename string
	Accepted string
}{
	ID:       "id",
	Wins:     "wins",
	Battles:  "battles",
	Name:     "name",
	Token:    "token",
	Filename: "filename",
	Accepted: "accepted",
}

var CardTableColumns = struct {
	ID       string
	Wins     string
	Battles  string
	Name     string
	Token    string
	Filename string
	Accepted string
}{
	ID:       "cards.id",
	Wins:     "cards.wins",
	Battles:  "cards.battles",
	Name:     "cards.name",
	Token:    "cards.token",
	Filename: "cards.filename",
	Accepted: "cards.accepted",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var CardWhere = struct {
	ID       whereHelperint
	Wins     whereHelperint
	Battles  whereHelperint
	Name     whereHelperstring
	Token    whereHelperstring
	Filename whereHelperstring
	Accepted whereHelperbool
}{
	ID:       whereHelperint{field: "\"cards\".\"id\""},
	Wins:     whereHelperint{field: "\"cards\".\"wins\""},
	Battles:  whereHelperint{field: "\"cards\".\"battles\""},
	Name:     whereHelperstring{field: "\"cards\".\"name\""},
	Token:    whereHelperstring{field: "\"cards\".\"token\""},
	Filename: whereHelperstring{field: "\"cards\".\"filename\""},
	Accepted: whereHelperbool{field: "\"cards\".\"accepted\""},
}

// CardRels is where relationship names are stored.
var CardRels = struct {
	Card1Battles string
	Card2Battles string
}{
	Card1Battles: "Card1Battles",
	Card2Battles: "Card2Battles",
}

// cardR is where relationships are stored.
type cardR struct {
	Card1Battles BattleSlice `boil:"Card1Battles" json:"Card1Battles" toml:"Card1Battles" yaml:"Card1Battles"`
	Card2Battles BattleSlice `boil:"Card2Battles" json:"Card2Battles" toml:"Card2Battles" yaml:"Card2Battles"`
}

// NewStruct creates a new relationship struct
func (*cardR) NewStruct() *cardR {
	return &cardR{}
}

func (r *cardR) GetCard1Battles() BattleSlice {
	if r == nil {
		return nil
	}
	return r.Card1Battles
}

func (r *cardR) GetCard2Battles() BattleSlice {
	if r == nil {
		return nil
	}
	return r.Card2Battles
}

// cardL is where Load methods for each relationship are stored.
type cardL struct{}

var (
	cardAllColumns            = []string{"id", "wins", "battles", "name", "token", "filename", "accepted"}
	cardColumnsWithoutDefault = []string{"name", "token", "filename"}
	cardColumnsWithDefault    = []string{"id", "wins", "battles", "accepted"}
	cardPrimaryKeyColumns     = []string{"id"}
	cardGeneratedColumns      = []string{}
)

type (
	// CardSlice is an alias for a slice of pointers to Card.
	// This should almost always be used instead of []Card.
	CardSlice []*Card
	// CardHook is the signature for custom Card hook methods
	CardHook func(context.Context, boil.ContextExecutor, *Card) error

	cardQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cardType                 = reflect.TypeOf(&Card{})
	cardMapping              = queries.MakeStructMapping(cardType)
	cardPrimaryKeyMapping, _ = queries.BindMapping(cardType, cardMapping, cardPrimaryKeyColumns)
	cardInsertCacheMut       sync.RWMutex
	cardInsertCache          = make(map[string]insertCache)
	cardUpdateCacheMut       sync.RWMutex
	cardUpdateCache          = make(map[string]updateCache)
	cardUpsertCacheMut       sync.RWMutex
	cardUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cardAfterSelectHooks []CardHook

var cardBeforeInsertHooks []CardHook
var cardAfterInsertHooks []CardHook

var cardBeforeUpdateHooks []CardHook
var cardAfterUpdateHooks []CardHook

var cardBeforeDeleteHooks []CardHook
var cardAfterDeleteHooks []CardHook

var cardBeforeUpsertHooks []CardHook
var cardAfterUpsertHooks []CardHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Card) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Card) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Card) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Card) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Card) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Card) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Card) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Card) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Card) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCardHook registers your hook function for all future operations.
func AddCardHook(hookPoint boil.HookPoint, cardHook CardHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		cardAfterSelectHooks = append(cardAfterSelectHooks, cardHook)
	case boil.BeforeInsertHook:
		cardBeforeInsertHooks = append(cardBeforeInsertHooks, cardHook)
	case boil.AfterInsertHook:
		cardAfterInsertHooks = append(cardAfterInsertHooks, cardHook)
	case boil.BeforeUpdateHook:
		cardBeforeUpdateHooks = append(cardBeforeUpdateHooks, cardHook)
	case boil.AfterUpdateHook:
		cardAfterUpdateHooks = append(cardAfterUpdateHooks, cardHook)
	case boil.BeforeDeleteHook:
		cardBeforeDeleteHooks = append(cardBeforeDeleteHooks, cardHook)
	case boil.AfterDeleteHook:
		cardAfterDeleteHooks = append(cardAfterDeleteHooks, cardHook)
	case boil.BeforeUpsertHook:
		cardBeforeUpsertHooks = append(cardBeforeUpsertHooks, cardHook)
	case boil.AfterUpsertHook:
		cardAfterUpsertHooks = append(cardAfterUpsertHooks, cardHook)
	}
}

// One returns a single card record from the query.
func (q cardQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Card, error) {
	o := &Card{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cards")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Card records from the query.
func (q cardQuery) All(ctx context.Context, exec boil.ContextExecutor) (CardSlice, error) {
	var o []*Card

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Card slice")
	}

	if len(cardAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Card records in the query.
func (q cardQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cards rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cardQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cards exists")
	}

	return count > 0, nil
}

// Card1Battles retrieves all the battle's Battles with an executor via card1_id column.
func (o *Card) Card1Battles(mods ...qm.QueryMod) battleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"battles\".\"card1_id\"=?", o.ID),
	)

	return Battles(queryMods...)
}

// Card2Battles retrieves all the battle's Battles with an executor via card2_id column.
func (o *Card) Card2Battles(mods ...qm.QueryMod) battleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"battles\".\"card2_id\"=?", o.ID),
	)

	return Battles(queryMods...)
}

// LoadCard1Battles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (cardL) LoadCard1Battles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCard interface{}, mods queries.Applicator) error {
	var slice []*Card
	var object *Card

	if singular {
		var ok bool
		object, ok = maybeCard.(*Card)
		if !ok {
			object = new(Card)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCard)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCard))
			}
		}
	} else {
		s, ok := maybeCard.(*[]*Card)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCard)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCard))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &cardR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cardR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`battles`),
		qm.WhereIn(`battles.card1_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load battles")
	}

	var resultSlice []*Battle
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice battles")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on battles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for battles")
	}

	if len(battleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Card1Battles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &battleR{}
			}
			foreign.R.Card1 = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.Card1ID {
				local.R.Card1Battles = append(local.R.Card1Battles, foreign)
				if foreign.R == nil {
					foreign.R = &battleR{}
				}
				foreign.R.Card1 = local
				break
			}
		}
	}

	return nil
}

// LoadCard2Battles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (cardL) LoadCard2Battles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCard interface{}, mods queries.Applicator) error {
	var slice []*Card
	var object *Card

	if singular {
		var ok bool
		object, ok = maybeCard.(*Card)
		if !ok {
			object = new(Card)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCard)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCard))
			}
		}
	} else {
		s, ok := maybeCard.(*[]*Card)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCard)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCard))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &cardR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cardR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`battles`),
		qm.WhereIn(`battles.card2_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load battles")
	}

	var resultSlice []*Battle
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice battles")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on battles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for battles")
	}

	if len(battleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Card2Battles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &battleR{}
			}
			foreign.R.Card2 = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.Card2ID {
				local.R.Card2Battles = append(local.R.Card2Battles, foreign)
				if foreign.R == nil {
					foreign.R = &battleR{}
				}
				foreign.R.Card2 = local
				break
			}
		}
	}

	return nil
}

// AddCard1Battles adds the given related objects to the existing relationships
// of the card, optionally inserting them as new records.
// Appends related to o.R.Card1Battles.
// Sets related.R.Card1 appropriately.
func (o *Card) AddCard1Battles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Battle) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.Card1ID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"battles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"card1_id"}),
				strmangle.WhereClause("\"", "\"", 2, battlePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.Card1ID = o.ID
		}
	}

	if o.R == nil {
		o.R = &cardR{
			Card1Battles: related,
		}
	} else {
		o.R.Card1Battles = append(o.R.Card1Battles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &battleR{
				Card1: o,
			}
		} else {
			rel.R.Card1 = o
		}
	}
	return nil
}

// AddCard2Battles adds the given related objects to the existing relationships
// of the card, optionally inserting them as new records.
// Appends related to o.R.Card2Battles.
// Sets related.R.Card2 appropriately.
func (o *Card) AddCard2Battles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Battle) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.Card2ID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"battles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"card2_id"}),
				strmangle.WhereClause("\"", "\"", 2, battlePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.Card2ID = o.ID
		}
	}

	if o.R == nil {
		o.R = &cardR{
			Card2Battles: related,
		}
	} else {
		o.R.Card2Battles = append(o.R.Card2Battles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &battleR{
				Card2: o,
			}
		} else {
			rel.R.Card2 = o
		}
	}
	return nil
}

// Cards retrieves all the records using an executor.
func Cards(mods ...qm.QueryMod) cardQuery {
	mods = append(mods, qm.From("\"cards\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"cards\".*"})
	}

	return cardQuery{q}
}

// FindCard retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCard(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Card, error) {
	cardObj := &Card{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cards\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cardObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cards")
	}

	if err = cardObj.doAfterSelectHooks(ctx, exec); err != nil {
		return cardObj, err
	}

	return cardObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Card) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cards provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cardColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cardInsertCacheMut.RLock()
	cache, cached := cardInsertCache[key]
	cardInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cardAllColumns,
			cardColumnsWithDefault,
			cardColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cardType, cardMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cardType, cardMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"cards\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"cards\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into cards")
	}

	if !cached {
		cardInsertCacheMut.Lock()
		cardInsertCache[key] = cache
		cardInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Card.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Card) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cardUpdateCacheMut.RLock()
	cache, cached := cardUpdateCache[key]
	cardUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cardAllColumns,
			cardPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cards, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cards\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cardPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cardType, cardMapping, append(wl, cardPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cards row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cards")
	}

	if !cached {
		cardUpdateCacheMut.Lock()
		cardUpdateCache[key] = cache
		cardUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cardQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cards")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CardSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"cards\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, cardPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in card slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all card")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Card) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cards provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cardColumnsWithDefault, o)

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

	cardUpsertCacheMut.RLock()
	cache, cached := cardUpsertCache[key]
	cardUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cardAllColumns,
			cardColumnsWithDefault,
			cardColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			cardAllColumns,
			cardPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert cards, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cardPrimaryKeyColumns))
			copy(conflict, cardPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"cards\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(cardType, cardMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cardType, cardMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert cards")
	}

	if !cached {
		cardUpsertCacheMut.Lock()
		cardUpsertCache[key] = cache
		cardUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Card record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Card) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Card provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cardPrimaryKeyMapping)
	sql := "DELETE FROM \"cards\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cards")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cardQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cardQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cards")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CardSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cardBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"cards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cardPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from card slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cards")
	}

	if len(cardAfterDeleteHooks) != 0 {
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
func (o *Card) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCard(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CardSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CardSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"cards\".* FROM \"cards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cardPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CardSlice")
	}

	*o = slice

	return nil
}

// CardExists checks if the Card row exists.
func CardExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"cards\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cards exists")
	}

	return exists, nil
}

// Exists checks if the Card row exists.
func (o *Card) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CardExists(ctx, exec, o.ID)
}
