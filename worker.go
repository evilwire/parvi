package parvi

// Worker - an interface representing a client or processor that
// handles requests (operations).
//
// For example, if you want to define a pool of map/reduce processes
// then a worker may be a logical unit (consisting of several logical
// units) that processes a request. If you want to define a pool of
// database clients, then a worker may be some struct that accepts a
// database query for an op, and returns the seeker for the row.
//
// Such a client/processor should know how to handle errors and clean
// up. The main function here would be its Process function which
// processes a certain operation and returns if it failed and if the
// operation should be retried.
type Worker interface {

	// Init - starts up a worker, returning any error that occurred while
	// starting up. If you do not need any initialisation, the Init function
	// can return nil.
	Init() error

	// Process - processes an operation. Returns whether or not the process
	// is in error and whether it should be retried. This is necessarily a
	// blocking operation, and should block until the process is deemed
	// complete.
	//
	// If you currently have a non-blocking processing function, then you
	// can adapt it here via a sync.WaitGroup.
	Process(op Operation) (bool, error)

	// HandleError - once the worker has ceased to
	HandleError(error, Operation)

	// Equal - determines whether a worker is the same as another. The
	// assumption here is that equality should be consistent but not
	// necessarily symmetric. That is, `a.Equal(a)` is true, and
	// `a.Equal(b)` is always the same regardless of the state of `b`.
	//
	// However, it is not required for `a.Equal(b)` to imply anything
	// about `b.Equal(a)`, e.g. even if `a.Equal(b)` is true, we may not
	// be able to conclude anything about `b.Equal(a)`.
	Equal(Worker) bool

	// Cleanup - cleans up any resource for the worker when the pool
	// shuts down
	Cleanup()
}
