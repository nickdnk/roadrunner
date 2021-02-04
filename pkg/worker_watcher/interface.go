package worker_watcher //nolint:golint,stylecheck

import (
	"context"

	"github.com/spiral/roadrunner/v2/pkg/worker"
)

// Watcher is an interface for the Sync workers lifecycle
type Watcher interface {
	// Watch used to add workers to the stack
	Watch(workers []worker.SyncWorker) error

	// Get provide first free worker
	Get(ctx context.Context) (worker.SyncWorker, error)

	// Push enqueues worker back
	Push(w worker.SyncWorker)

	// Allocate - allocates new worker and put it into the WorkerWatcher
	Allocate() error

	// Destroy destroys the underlying stack
	Destroy(ctx context.Context)

	// WorkersList return all stack w/o removing it from internal storage
	List() []worker.SyncWorker

	// RemoveWorker remove worker from the stack
	Remove(wb worker.SyncWorker) error
}
