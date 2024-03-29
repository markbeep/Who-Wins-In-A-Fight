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

// Battle is an object representing the database table.
type Battle struct {
	ID        int `boil:"id" json:"id" toml:"id" yaml:"id"`
	Card1ID   int `boil:"card1_id" json:"card1_id" toml:"card1_id" yaml:"card1_id"`
	Card2ID   int `boil:"card2_id" json:"card2_id" toml:"card2_id" yaml:"card2_id"`
	Card1Wins int `boil:"card1_wins" json:"card1_wins" toml:"card1_wins" yaml:"card1_wins"`
	Card2Wins int `boil:"card2_wins" json:"card2_wins" toml:"card2_wins" yaml:"card2_wins"`

	R *battleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L battleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BattleColumns = struct {
	ID        string
	Card1ID   string
	Card2ID   string
	Card1Wins string
	Card2Wins string
}{
	ID:        "id",
	Card1ID:   "card1_id",
	Card2ID:   "card2_id",
	Card1Wins: "card1_wins",
	Card2Wins: "card2_wins",
}

var BattleTableColumns = struct {
	ID        string
	Card1ID   string
	Card2ID   string
	Card1Wins string
	Card2Wins string
}{
	ID:        "battles.id",
	Card1ID:   "battles.card1_id",
	Card2ID:   "battles.card2_id",
	Card1Wins: "battles.card1_wins",
	Card2Wins: "battles.card2_wins",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BattleWhere = struct {
	ID        whereHelperint
	Card1ID   whereHelperint
	Card2ID   whereHelperint
	Card1Wins whereHelperint
	Card2Wins whereHelperint
}{
	ID:        whereHelperint{field: "\"battles\".\"id\""},
	Card1ID:   whereHelperint{field: "\"battles\".\"card1_id\""},
	Card2ID:   whereHelperint{field: "\"battles\".\"card2_id\""},
	Card1Wins: whereHelperint{field: "\"battles\".\"card1_wins\""},
	Card2Wins: whereHelperint{field: "\"battles\".\"card2_wins\""},
}

// BattleRels is where relationship names are stored.
var BattleRels = struct {
	Card1 string
	Card2 string
}{
	Card1: "Card1",
	Card2: "Card2",
}

// battleR is where relationships are stored.
type battleR struct {
	Card1 *Card `boil:"Card1" json:"Card1" toml:"Card1" yaml:"Card1"`
	Card2 *Card `boil:"Card2" json:"Card2" toml:"Card2" yaml:"Card2"`
}

// NewStruct creates a new relationship struct
func (*battleR) NewStruct() *battleR {
	return &battleR{}
}

func (r *battleR) GetCard1() *Card {
	if r == nil {
		return nil
	}
	return r.Card1
}

func (r *battleR) GetCard2() *Card {
	if r == nil {
		return nil
	}
	return r.Card2
}

// battleL is where Load methods for each relationship are stored.
type battleL struct{}

var (
	battleAllColumns            = []string{"id", "card1_id", "card2_id", "card1_wins", "card2_wins"}
	battleColumnsWithoutDefault = []string{"card1_id", "card2_id"}
	battleColumnsWithDefault    = []string{"id", "card1_wins", "card2_wins"}
	battlePrimaryKeyColumns     = []string{"id"}
	battleGeneratedColumns      = []string{}
)

type (
	// BattleSlice is an alias for a slice of pointers to Battle.
	// This should almost always be used instead of []Battle.
	BattleSlice []*Battle
	// BattleHook is the signature for custom Battle hook methods
	BattleHook func(context.Context, boil.ContextExecutor, *Battle) error

	battleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	battleType                 = reflect.TypeOf(&Battle{})
	battleMapping              = queries.MakeStructMapping(battleType)
	battlePrimaryKeyMapping, _ = queries.BindMapping(battleType, battleMapping, battlePrimaryKeyColumns)
	battleInsertCacheMut       sync.RWMutex
	battleInsertCache          = make(map[string]insertCache)
	battleUpdateCacheMut       sync.RWMutex
	battleUpdateCache          = make(map[string]updateCache)
	battleUpsertCacheMut       sync.RWMutex
	battleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var battleAfterSelectHooks []BattleHook

var battleBeforeInsertHooks []BattleHook
var battleAfterInsertHooks []BattleHook

var battleBeforeUpdateHooks []BattleHook
var battleAfterUpdateHooks []BattleHook

var battleBeforeDeleteHooks []BattleHook
var battleAfterDeleteHooks []BattleHook

var battleBeforeUpsertHooks []BattleHook
var battleAfterUpsertHooks []BattleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Battle) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Battle) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Battle) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Battle) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Battle) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Battle) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Battle) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Battle) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Battle) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range battleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBattleHook registers your hook function for all future operations.
func AddBattleHook(hookPoint boil.HookPoint, battleHook BattleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		battleAfterSelectHooks = append(battleAfterSelectHooks, battleHook)
	case boil.BeforeInsertHook:
		battleBeforeInsertHooks = append(battleBeforeInsertHooks, battleHook)
	case boil.AfterInsertHook:
		battleAfterInsertHooks = append(battleAfterInsertHooks, battleHook)
	case boil.BeforeUpdateHook:
		battleBeforeUpdateHooks = append(battleBeforeUpdateHooks, battleHook)
	case boil.AfterUpdateHook:
		battleAfterUpdateHooks = append(battleAfterUpdateHooks, battleHook)
	case boil.BeforeDeleteHook:
		battleBeforeDeleteHooks = append(battleBeforeDeleteHooks, battleHook)
	case boil.AfterDeleteHook:
		battleAfterDeleteHooks = append(battleAfterDeleteHooks, battleHook)
	case boil.BeforeUpsertHook:
		battleBeforeUpsertHooks = append(battleBeforeUpsertHooks, battleHook)
	case boil.AfterUpsertHook:
		battleAfterUpsertHooks = append(battleAfterUpsertHooks, battleHook)
	}
}

// One returns a single battle record from the query.
func (q battleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Battle, error) {
	o := &Battle{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for battles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Battle records from the query.
func (q battleQuery) All(ctx context.Context, exec boil.ContextExecutor) (BattleSlice, error) {
	var o []*Battle

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Battle slice")
	}

	if len(battleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Battle records in the query.
func (q battleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count battles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q battleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if battles exists")
	}

	return count > 0, nil
}

// Card1 pointed to by the foreign key.
func (o *Battle) Card1(mods ...qm.QueryMod) cardQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.Card1ID),
	}

	queryMods = append(queryMods, mods...)

	return Cards(queryMods...)
}

// Card2 pointed to by the foreign key.
func (o *Battle) Card2(mods ...qm.QueryMod) cardQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.Card2ID),
	}

	queryMods = append(queryMods, mods...)

	return Cards(queryMods...)
}

// LoadCard1 allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (battleL) LoadCard1(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBattle interface{}, mods queries.Applicator) error {
	var slice []*Battle
	var object *Battle

	if singular {
		var ok bool
		object, ok = maybeBattle.(*Battle)
		if !ok {
			object = new(Battle)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBattle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBattle))
			}
		}
	} else {
		s, ok := maybeBattle.(*[]*Battle)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBattle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBattle))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &battleR{}
		}
		args = append(args, object.Card1ID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &battleR{}
			}

			for _, a := range args {
				if a == obj.Card1ID {
					continue Outer
				}
			}

			args = append(args, obj.Card1ID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`cards`),
		qm.WhereIn(`cards.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Card")
	}

	var resultSlice []*Card
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Card")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for cards")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for cards")
	}

	if len(cardAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Card1 = foreign
		if foreign.R == nil {
			foreign.R = &cardR{}
		}
		foreign.R.Card1Battles = append(foreign.R.Card1Battles, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Card1ID == foreign.ID {
				local.R.Card1 = foreign
				if foreign.R == nil {
					foreign.R = &cardR{}
				}
				foreign.R.Card1Battles = append(foreign.R.Card1Battles, local)
				break
			}
		}
	}

	return nil
}

// LoadCard2 allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (battleL) LoadCard2(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBattle interface{}, mods queries.Applicator) error {
	var slice []*Battle
	var object *Battle

	if singular {
		var ok bool
		object, ok = maybeBattle.(*Battle)
		if !ok {
			object = new(Battle)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBattle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBattle))
			}
		}
	} else {
		s, ok := maybeBattle.(*[]*Battle)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBattle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBattle))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &battleR{}
		}
		args = append(args, object.Card2ID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &battleR{}
			}

			for _, a := range args {
				if a == obj.Card2ID {
					continue Outer
				}
			}

			args = append(args, obj.Card2ID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`cards`),
		qm.WhereIn(`cards.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Card")
	}

	var resultSlice []*Card
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Card")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for cards")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for cards")
	}

	if len(cardAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Card2 = foreign
		if foreign.R == nil {
			foreign.R = &cardR{}
		}
		foreign.R.Card2Battles = append(foreign.R.Card2Battles, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Card2ID == foreign.ID {
				local.R.Card2 = foreign
				if foreign.R == nil {
					foreign.R = &cardR{}
				}
				foreign.R.Card2Battles = append(foreign.R.Card2Battles, local)
				break
			}
		}
	}

	return nil
}

// SetCard1 of the battle to the related item.
// Sets o.R.Card1 to related.
// Adds o to related.R.Card1Battles.
func (o *Battle) SetCard1(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Card) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"battles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"card1_id"}),
		strmangle.WhereClause("\"", "\"", 2, battlePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Card1ID = related.ID
	if o.R == nil {
		o.R = &battleR{
			Card1: related,
		}
	} else {
		o.R.Card1 = related
	}

	if related.R == nil {
		related.R = &cardR{
			Card1Battles: BattleSlice{o},
		}
	} else {
		related.R.Card1Battles = append(related.R.Card1Battles, o)
	}

	return nil
}

// SetCard2 of the battle to the related item.
// Sets o.R.Card2 to related.
// Adds o to related.R.Card2Battles.
func (o *Battle) SetCard2(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Card) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"battles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"card2_id"}),
		strmangle.WhereClause("\"", "\"", 2, battlePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Card2ID = related.ID
	if o.R == nil {
		o.R = &battleR{
			Card2: related,
		}
	} else {
		o.R.Card2 = related
	}

	if related.R == nil {
		related.R = &cardR{
			Card2Battles: BattleSlice{o},
		}
	} else {
		related.R.Card2Battles = append(related.R.Card2Battles, o)
	}

	return nil
}

// Battles retrieves all the records using an executor.
func Battles(mods ...qm.QueryMod) battleQuery {
	mods = append(mods, qm.From("\"battles\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"battles\".*"})
	}

	return battleQuery{q}
}

// FindBattle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBattle(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Battle, error) {
	battleObj := &Battle{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"battles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, battleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from battles")
	}

	if err = battleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return battleObj, err
	}

	return battleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Battle) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no battles provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(battleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	battleInsertCacheMut.RLock()
	cache, cached := battleInsertCache[key]
	battleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			battleAllColumns,
			battleColumnsWithDefault,
			battleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(battleType, battleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(battleType, battleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"battles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"battles\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into battles")
	}

	if !cached {
		battleInsertCacheMut.Lock()
		battleInsertCache[key] = cache
		battleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Battle.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Battle) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	battleUpdateCacheMut.RLock()
	cache, cached := battleUpdateCache[key]
	battleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			battleAllColumns,
			battlePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update battles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"battles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, battlePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(battleType, battleMapping, append(wl, battlePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update battles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for battles")
	}

	if !cached {
		battleUpdateCacheMut.Lock()
		battleUpdateCache[key] = cache
		battleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q battleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for battles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for battles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BattleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), battlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"battles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, battlePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in battle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all battle")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Battle) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no battles provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(battleColumnsWithDefault, o)

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

	battleUpsertCacheMut.RLock()
	cache, cached := battleUpsertCache[key]
	battleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			battleAllColumns,
			battleColumnsWithDefault,
			battleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			battleAllColumns,
			battlePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert battles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(battlePrimaryKeyColumns))
			copy(conflict, battlePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"battles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(battleType, battleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(battleType, battleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert battles")
	}

	if !cached {
		battleUpsertCacheMut.Lock()
		battleUpsertCache[key] = cache
		battleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Battle record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Battle) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Battle provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), battlePrimaryKeyMapping)
	sql := "DELETE FROM \"battles\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from battles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for battles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q battleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no battleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from battles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for battles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BattleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(battleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), battlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"battles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, battlePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from battle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for battles")
	}

	if len(battleAfterDeleteHooks) != 0 {
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
func (o *Battle) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBattle(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BattleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BattleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), battlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"battles\".* FROM \"battles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, battlePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BattleSlice")
	}

	*o = slice

	return nil
}

// BattleExists checks if the Battle row exists.
func BattleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"battles\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if battles exists")
	}

	return exists, nil
}

// Exists checks if the Battle row exists.
func (o *Battle) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BattleExists(ctx, exec, o.ID)
}
