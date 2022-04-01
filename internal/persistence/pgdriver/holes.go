package pgdriver

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/nkralles/masters-web/internal/persistence"
)

func (d *Driver) GetHole(ctx context.Context, hole int) (*persistence.Hole, error) {
	var h persistence.Hole
	err := d.pool.QueryRow(ctx, `
			select hole_number, name, par, yards
			from masters_holes
			where hole_number = $1;
			`, hole).Scan(&h.Hole, &h.Name, &h.Par, &h.Yards)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, persistence.ErrHoleNotFound(hole)
		}
		return nil, err
	}
	return &h, nil
}

func (d *Driver) GetHoles(ctx context.Context) (*[]persistence.Hole, error) {
	rows, err := d.pool.Query(ctx, `select hole_number, name, par, yards
							from masters_holes
							order by hole_number;`)

	if err != nil {
		return nil, err
	}
	holes := make([]persistence.Hole, 0)
	for rows.Next() {
		var hole persistence.Hole
		if err = rows.Scan(&hole.Hole, &hole.Name, &hole.Par, &hole.Yards); err != nil {
			return nil, err
		}
		holes = append(holes, hole)
	}
	return &holes, nil
}
