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

func (d *Driver) GetScores(ctx context.Context) ([]persistence.Score, error) {
	scores := make([]persistence.Score, 0)

	rows, err := d.pool.Query(ctx,
		`select *, rank() over (order by total) as standing
from (
         select t1.player_id,
                rank,
                first_name,
                last_name,
                cc,
                jsonb_agg(jsonb_build_object('round', round, 'toPar', score, 'lastupdated', ts)) as rounds,
                case when sum(score) > 1000 then 1000 else sum(score) end                        as total,
                max(ts)                                                                          as last_updated
         from (
                  select distinct on (player_id, round) player_id, score, round, ts
                  from masters_scores
                  order by player_id, round, ts desc
              ) t1
                  join golfers g on t1.player_id = g.player_id
         group by t1.player_id, rank, first_name, last_name, cc
         order by total
     ) t2;`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s persistence.Score
		err = rows.Scan(&s.PlayerId, &s.Rank, &s.FirstName, &s.LastName, &s.CountryCode, &s.Rounds, &s.ToPar, &s.LastUpdated, &s.Standing)
		if err != nil {
			return nil, err
		}
		scores = append(scores, s)
	}
	return scores, nil
}
