package persistence

import "fmt"

var (
	ErrHoleNotFound      = func(hole int) error { return fmt.Errorf("hole %d not found", hole) }
	ErrGolferNotFound    = func(pid int) error { return fmt.Errorf("golfer with player id %d not found", pid) }
	ErrGolferNotFoundStr = func(name string) error { return fmt.Errorf("golfer %s not found", name) }
)
