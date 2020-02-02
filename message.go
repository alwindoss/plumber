package plumber

import "io"

// Message is an interface that needs to be implemented by the messages of the pipeline
type Message interface {
	io.ReadWriter
}
