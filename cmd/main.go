package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
)

var version = "0.0.1-dev"

func catfile(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)
}

func main() {

	fmt.Println("B+ Monitor ", version)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Listen for events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("created file: ", event.Name)
					catfile(event.Name)
				}

				// log.Println("event: ", event)
				// if event.Has(fsnotify.Create) {
				// 	log.Println("created file: ", event.Name)
				// }
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.Add("./content")
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}
