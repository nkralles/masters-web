package persistence

import "context"

var defaultDriver MastersStorage = nil

type MastersStorage interface {
	GetHoles(ctx context.Context) (*[]Hole, error)
}

func DefaultDriver() MastersStorage {
	return defaultDriver
}

func SetDefaultDriver(driver MastersStorage) {
	defaultDriver = driver

}
