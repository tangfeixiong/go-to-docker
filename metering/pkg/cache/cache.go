package cache

import info "github.com/google/cadvisor/info/v1"

type Cache interface {
	// Add a ContainerStats for the specified container.
	AddStats(ref info.ContainerReference, stats *info.ContainerStats) error

	// Remove all cached information for the specified container.
	RemoveContainer(containerName string) error

	// Read most recent stats. numStats indicates max number of stats
	// returned. The returned stats must be consecutive observed stats. If
	// numStats < 0, then return all stats stored in the storage. The
	// returned stats should be sorted in time increasing order, i.e. Most
	// recent stats should be the last.
	RecentStats(containerName string, numStats int) ([]*info.ContainerStats, error)

	// Close will clear the state of the storage driver. The elements
	// stored in the underlying storage may or may not be deleted depending
	// on the implementation of the storage driver.
	Close() error
}
