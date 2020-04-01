package filewatcher

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type FileWatcher struct {
	Watcher *fsnotify.Watcher
}

func NewFileWatcher() *FileWatcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	return &FileWatcher{
		Watcher: watcher,
	}
}

func (f *FileWatcher) Watch(directory string) {
	log.Println("Starting the watch .. ")
	fmt.Println("watching : ", directory+"/test/foo")
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-f.Watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-f.Watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err := f.Watcher.Add(directory + "/test/foo")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
