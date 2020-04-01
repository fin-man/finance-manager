package filewatcher

import (
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

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-f.Watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
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

	err := f.Watcher.Add("/test/foo")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}