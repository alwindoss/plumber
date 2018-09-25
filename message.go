package plumber

// type Message1 interface {
// 	Get() []byte
// 	Set([]byte)
// }

// Message is an interface that needs to be implemented by the messages of the pipeline
type Message struct {
	Header []byte
	Body   []byte
}
