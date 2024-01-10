package storage

import (
	"sync"
	"time"
)

type BattleCard struct {
	Url    string
	ID     int
	Wins   int
	Losses int
	Name   string
}

type Battle struct {
	Start         time.Time
	Card1         BattleCard
	Card2         BattleCard
	Token         string
	CategoryToken string
}

type Category struct {
	Token              string
	Title              string
	AllCards           map[int]*BattleCard
	AllCardsMutex      sync.RWMutex
	ActiveBattles      map[string]Battle
	ActiveBattlesMutex sync.RWMutex
}
