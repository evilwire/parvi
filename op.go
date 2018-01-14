package parvi


// Operation - an interface representing a unit of task to be
// done by a worker. For example, if you want to define a pool of
// http requests or database readers, then your operation can be
// a request object or an object that represents a database read.
// To capture the output, either publicly or privately set an output.
//
// The interface assumes that Wait() would block until the operation
// is Done(). Otherwise, the operation may be a bit unpredictable.
type Operation interface {
	// IncrementTry - increments a counter that the operation is
	// tried again
	IncrementTry()

	// Wait - Blocks the flow of execution till the operation is done.
	// Must be able to call Wait() from various points in code.
	Wait()

	// Done - signals that the Operation is done. Anyone waiting
	// with Wait() will resume.
	//
	// Unfortunate for you, this has to be idempotent.
	// Minimally, it has to not crash when called consecutively
	// as we don't always guarantee exactly-once execution, and
	// definitely not in invoking this method
	Done()
}