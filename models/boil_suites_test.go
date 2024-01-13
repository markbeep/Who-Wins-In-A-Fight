// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Battles", testBattles)
	t.Run("Cards", testCards)
	t.Run("UserTokens", testUserTokens)
}

func TestDelete(t *testing.T) {
	t.Run("Battles", testBattlesDelete)
	t.Run("Cards", testCardsDelete)
	t.Run("UserTokens", testUserTokensDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Battles", testBattlesQueryDeleteAll)
	t.Run("Cards", testCardsQueryDeleteAll)
	t.Run("UserTokens", testUserTokensQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Battles", testBattlesSliceDeleteAll)
	t.Run("Cards", testCardsSliceDeleteAll)
	t.Run("UserTokens", testUserTokensSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Battles", testBattlesExists)
	t.Run("Cards", testCardsExists)
	t.Run("UserTokens", testUserTokensExists)
}

func TestFind(t *testing.T) {
	t.Run("Battles", testBattlesFind)
	t.Run("Cards", testCardsFind)
	t.Run("UserTokens", testUserTokensFind)
}

func TestBind(t *testing.T) {
	t.Run("Battles", testBattlesBind)
	t.Run("Cards", testCardsBind)
	t.Run("UserTokens", testUserTokensBind)
}

func TestOne(t *testing.T) {
	t.Run("Battles", testBattlesOne)
	t.Run("Cards", testCardsOne)
	t.Run("UserTokens", testUserTokensOne)
}

func TestAll(t *testing.T) {
	t.Run("Battles", testBattlesAll)
	t.Run("Cards", testCardsAll)
	t.Run("UserTokens", testUserTokensAll)
}

func TestCount(t *testing.T) {
	t.Run("Battles", testBattlesCount)
	t.Run("Cards", testCardsCount)
	t.Run("UserTokens", testUserTokensCount)
}

func TestHooks(t *testing.T) {
	t.Run("Battles", testBattlesHooks)
	t.Run("Cards", testCardsHooks)
	t.Run("UserTokens", testUserTokensHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Battles", testBattlesInsert)
	t.Run("Battles", testBattlesInsertWhitelist)
	t.Run("Cards", testCardsInsert)
	t.Run("Cards", testCardsInsertWhitelist)
	t.Run("UserTokens", testUserTokensInsert)
	t.Run("UserTokens", testUserTokensInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BattleToCardUsingCard1", testBattleToOneCardUsingCard1)
	t.Run("BattleToCardUsingCard2", testBattleToOneCardUsingCard2)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CardToCard1Battles", testCardToManyCard1Battles)
	t.Run("CardToCard2Battles", testCardToManyCard2Battles)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BattleToCardUsingCard1Battles", testBattleToOneSetOpCardUsingCard1)
	t.Run("BattleToCardUsingCard2Battles", testBattleToOneSetOpCardUsingCard2)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CardToCard1Battles", testCardToManyAddOpCard1Battles)
	t.Run("CardToCard2Battles", testCardToManyAddOpCard2Battles)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Battles", testBattlesReload)
	t.Run("Cards", testCardsReload)
	t.Run("UserTokens", testUserTokensReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Battles", testBattlesReloadAll)
	t.Run("Cards", testCardsReloadAll)
	t.Run("UserTokens", testUserTokensReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Battles", testBattlesSelect)
	t.Run("Cards", testCardsSelect)
	t.Run("UserTokens", testUserTokensSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Battles", testBattlesUpdate)
	t.Run("Cards", testCardsUpdate)
	t.Run("UserTokens", testUserTokensUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Battles", testBattlesSliceUpdateAll)
	t.Run("Cards", testCardsSliceUpdateAll)
	t.Run("UserTokens", testUserTokensSliceUpdateAll)
}