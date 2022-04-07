package pgdriver

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/nkralles/masters-web/internal/persistence"
)

func (d *Driver) AddGolfer(ctx context.Context, golfer *persistence.Golfer) error {
	_, err := d.pool.Exec(ctx, `insert into golfers(player_id, rank, first_name, last_name, cc) values ($1, $2, $3, $4, $5) 
							on conflict (player_id) do update set rank=$2, first_name=$3, last_name=$4, cc=$5;`,
		golfer.PlayerId, golfer.Rank, golfer.FirstName, golfer.LastName, golfer.CountryCode)
	return err
}

func (d *Driver) GetGolfers(ctx context.Context, params *persistence.CommonParams) (*persistence.GolferResponse, error) {
	wheres := make([]string, 0)

	if params.TextParams != nil && len(params.Query) != 0 {
		wheres = append(wheres, ColumnLike(fmt.Sprintf("(first_name || ' ' || last_name)"), params.Query))
	}
	res := new(persistence.GolferResponse)
	res.Golfers = &[]persistence.Golfer{}

	whereStr := WhereAnds(wheres)
	err := d.pool.QueryRow(ctx, fmt.Sprintf(`select count(*) from golfers %s`,
		whereStr)).Scan(&res.Total)
	if err != nil {
		return nil, err
	}

	rows, err := d.pool.Query(ctx, fmt.Sprintf(`select player_id, rank, first_name, last_name, cc from golfers %s %s %s %s`,
		whereStr, OrderBy([]string{"rank"}, true), Limit64(params.Limit), Offset64(params.Offset)))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var golfer persistence.Golfer
		if err = rows.Scan(&golfer.PlayerId, &golfer.Rank, &golfer.FirstName, &golfer.LastName, &golfer.CountryCode); err != nil {
			return nil, err
		}
		*res.Golfers = append(*res.Golfers, golfer)
	}

	res.Count = int64(len(*res.Golfers))

	return res, nil

}

func (d *Driver) GetGolferById(ctx context.Context, id int) (*persistence.Golfer, error) {
	golfer := new(persistence.Golfer)
	err := d.pool.QueryRow(ctx, `select player_id, rank, first_name, last_name, cc from golfers where player_id = $1`,
		id).Scan(&golfer.PlayerId, &golfer.Rank, &golfer.FirstName, &golfer.LastName, &golfer.CountryCode)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, persistence.ErrGolferNotFound(id)
		}
		return nil, err
	}
	return golfer, nil
}

var golferMap = map[string]string{
	"Jose Maria Olazabal": "Jose M Olazabal",
	"J. J. Spaun":         "J.J. Spaun",
	"Si Woo Kim":          "SiWoo Kim",
	"James Piot":          "James Piot(Am)",
	"Laird Shepherd":      "Laird Shepherd(Am)",
	"Keita Nakajima":      "Keita Nakajima(Am)",
	"Matthew Fitzpatrick": "Matt Fitzpatrick",
}

func (d *Driver) GetGolferByFullName(ctx context.Context, name string) (*persistence.Golfer, error) {
	golfer := new(persistence.Golfer)

	if v, ok := golferMap[name]; ok {
		name = v
	}

	err := d.pool.QueryRow(ctx, `select player_id, rank, first_name, last_name, cc from golfers where (first_name || ' ' || last_name)::citext = $1`,
		name).Scan(&golfer.PlayerId, &golfer.Rank, &golfer.FirstName, &golfer.LastName, &golfer.CountryCode)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, persistence.ErrGolferNotFoundStr(name)
		}
		return nil, err
	}
	return golfer, nil
}
