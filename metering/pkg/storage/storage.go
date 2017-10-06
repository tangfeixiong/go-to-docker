package storage

import (
	"fmt"
	"sort"

	info "github.com/google/cadvisor/info/v1"
)

type StorageDriver interface {
	AddStats(ref info.ContainerReference, stats *info.ContainerStats) error

	// Close will clear the state of the storage driver. The elements
	// stored in the underlying storage may or may not be deleted depending
	// on the implementation of the storage driver.
	Close() error
}

type StorageDriverFunc func() (StorageDriver, error)

var registeredPlugins = map[string](StorageDriverFunc){}

func RegisterStorageDriver(name string, f StorageDriverFunc) {
	registeredPlugins[name] = f
}

func New(name string) (StorageDriver, error) {
	if name == "" {
		return nil, nil
	}
	f, ok := registeredPlugins[name]
	if !ok {
		return nil, fmt.Errorf("unknown backend storage driver: %s", name)
	}
	return f()
}

func ListDrivers() []string {
	drivers := make([]string, 0, len(registeredPlugins))
	for name := range registeredPlugins {
		drivers = append(drivers, name)
	}
	sort.Strings(drivers)
	return drivers
}
