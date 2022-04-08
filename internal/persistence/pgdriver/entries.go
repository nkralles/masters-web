package pgdriver

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/nkralles/masters-web/internal/persistence"
)

func (d *Driver) SetEntryWinningScore(ctx context.Context, name string, score int) error {
	entry, err := d.getEntryUser(ctx, name)
	if err != nil {
		return err
	}
	cmd, err := d.pool.Exec(ctx, `update entries set winning_score = $2 where name = $1;`, entry.Name, score)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("failed to update score for user %s", entry.Name)
	}
	return nil
}

func (d *Driver) GetEntries(ctx context.Context) (*[]persistence.Entry, error) {
	entries := make([]persistence.Entry, 0)
	rows, err := d.pool.Query(ctx, `
select *, row_number() over (order by total) as rank
from (select name,
             winning_score,
             golfers,
             (select sum(_s.s) from (select unnest(total) as s order by 1 limit 5) _s) as total
      from (select name,
                   winning_score,
                   coalesce(jsonb_agg(golfer order by (golfer ->> 'rank')::int), jsonb_build_array()) as golfers,
                   array_agg(coalesce(par, 0))                                                        as total
            from (select t1.name,
                         t1.winning_score,
                         jsonb_build_object('playerId', t1.player_id, 'rank', rank, 'first_name', first_name,
                                            'last_name',
                                            last_name, 'cc', cc, 'toPar', coalesce(total_score.score, 0)) as golfer,
                         coalesce(total_score.score, 0)                                                   as par
                  from (select name, winning_score, g.player_id, g.rank, g.first_name, g.last_name, g.cc
                        from entries e
                                 join user_golfer_entries uge on e.name = uge.entry_name
                                 join golfers g on g.player_id = uge.golfer_id) t1
                           left outer join (select distinct on (player_id) player_id, score
                                            from masters_scores
                                            order by player_id, ts desc) as total_score
                                           on total_score.player_id = t1.player_id
                  order by t1.name) t3
            group by name, winning_score
            order by total) t4) t5;
		`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var e persistence.Entry
		var golfers []map[string]any
		if err = rows.Scan(&e.Name, &e.WinningScore, &golfers, &e.Total, &e.Rank); err != nil {
			return nil, err
		}
		b, err := json.Marshal(golfers)
		if err != nil {
			return nil, err
		}
		enc := json.NewDecoder(bytes.NewReader(b))
		err = enc.Decode(&e.Golfers)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return &entries, nil
}

func (d *Driver) getEntryUser(ctx context.Context, name string) (*persistence.Entry, error) {
	entry := new(persistence.Entry)
	var golfers []map[string]any
	if err := d.pool.QueryRow(ctx, `
			select name,
       winning_score,
       coalesce(jsonb_agg(jsonb_build_object(
                                  'playerId', g.player_id,
                                  'rank', g.rank,
                                  'first_name', g.first_name,
                                  'last_name', g.last_name,
                                  'cc', g.cc)
                          order by g.rank), jsonb_build_array())
					from entries e
							 left outer join user_golfer_entries uge on e.name = uge.entry_name
							 left join golfers g on g.player_id = uge.golfer_id
					where e.name = $1
					group by name, winning_score;
		`,
		name).Scan(&entry.Name, &entry.WinningScore, &golfers); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = d.pool.QueryRow(ctx, `insert into entries (name) values($1) returning name,winning_score`,
				name).Scan(&entry.Name, &entry.WinningScore)
			if err != nil {
				return nil, err
			}
			return entry, nil
		}
		return nil, err
	}
	b, err := json.Marshal(golfers)
	if err != nil {
		return nil, err
	}
	enc := json.NewDecoder(bytes.NewReader(b))
	err = enc.Decode(&entry.Golfers)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (d *Driver) GetEntryUser(ctx context.Context, name string) (*persistence.Entry, error) {
	return d.getEntryUser(ctx, name)
}

func (d *Driver) DeleteAllEntriesForUser(ctx context.Context, entry *persistence.Entry) error {
	_, err := d.pool.Exec(ctx, `delete from user_golfer_entries where entry_name = $1`, entry.Name)
	return err
}

func (d *Driver) AddGolferEntryForUser(ctx context.Context, entry *persistence.Entry, golfer *persistence.Golfer) error {
	_, err := d.pool.Exec(ctx, `insert into user_golfer_entries(entry_name, golfer_id) VALUES ($1, $2)`,
		entry.Name, golfer.PlayerId)
	return err
}
