package pgdriver

import (
	"context"
	"github.com/nkralles/masters-web/internal/persistence"
)

func (d *Driver) AddScore(ctx context.Context, golfer *persistence.Golfer, round, score int) error {
	_, err := d.pool.Exec(ctx, `insert into masters_scores(player_id, score, round) VALUES ($1, $2, $3)`,
		golfer.PlayerId, score, round)
	return err
}
