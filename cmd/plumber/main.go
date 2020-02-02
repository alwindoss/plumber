package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/alwindoss/plumber"
)

func main() {
	fmt.Println("Welcome to plumber")
	src := plumber.NewFileSource("/opt/alwin/tmp/dummyfolder/")
	srcCh := make(chan plumber.Message, 10)
	go src.Source(srcCh)
	for d := range srcCh {
		time.Sleep(2 * time.Second)
		data, err := ioutil.ReadAll(d)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		log.Printf("MAIN::Data Received: %v", string(data))
	}
}
