package plumber

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Sourcer is the interface that defines the contract for source of a pipeline
type Sourcer interface {
	Source(out chan Message)
}

// NewFileSource creates a new source which fetches data from the filesystem
func NewFileSource(loc string) Sourcer {
	return &filesystemSource{
		loc: loc,
	}
}

type filesystemSource struct {
	loc string
}

func (f filesystemSource) Source(out chan Message) {
	defer close(out)
	filepath.Walk(f.loc, func(path string, info os.FileInfo, err error) error {
		log.Printf("SOURCE::path: %s", path)
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			err = fmt.Errorf("unable to open the file: %w", err)
			return err
		}
		out <- f
		return nil
	})
}
