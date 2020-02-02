package plumber

// Sinker is the interface that defines the contract for source of a pipeline
type Sinker interface {
	Sink(in chan Message)
}
