package plumber

// Sourcer is the interface that defines the contract for source of a pipeline
type Sourcer interface {
	Source(out chan *Message)
}
