package persistence

import (
	"context"
)

var defaultDriver MastersStorage = nil

type MastersStorage interface {
	AddGolfer(ctx context.Context, golfer *Golfer) error
	GetGolfers(ctx context.Context, params *CommonParams) (*GolferResponse, error)
	GetGolferById(ctx context.Context, id int) (*Golfer, error)
	GetGolferByFullName(ctx context.Context, name string) (*Golfer, error)

	GetHoles(ctx context.Context) (*[]Hole, error)
	GetHole(ctx context.Context, hole int) (*Hole, error)

	GetEntryUser(ctx context.Context, name string) (*Entry, error)
	SetEntryWinningScore(ctx context.Context, name string, score int) error
	DeleteAllEntriesForUser(ctx context.Context, entry *Entry) error
	AddGolferEntryForUser(ctx context.Context, entry *Entry, golfer *Golfer) error
	GetEntries(ctx context.Context) (*[]Entry, error)

	AddScore(ctx context.Context, golfer *Golfer, round, score int) error
	GetScores(ctx context.Context) ([]Score, error)
}

func DefaultDriver() MastersStorage {
	return defaultDriver
}

func SetDefaultDriver(driver MastersStorage) {
	defaultDriver = driver
}
