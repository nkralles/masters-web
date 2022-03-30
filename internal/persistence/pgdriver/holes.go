package pgdriver

import (
	"context"
	"github.com/nkralles/masters-web/internal/persistence"
)

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
