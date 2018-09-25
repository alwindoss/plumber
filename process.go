package plumber

// Processor defines the method that the type should implement if it has to be added to the pipeline
type Processor interface {
	Process(in, out chan *Message)
}
