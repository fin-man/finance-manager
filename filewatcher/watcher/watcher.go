package watcher
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

func (f *FileWatcher) Watch(volume string) {
	log.Println("Starting the watch .. ")
	fmt.Println("watching : ", volume)
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-f.Watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("created file:", event.Name)
				}
			case err, ok := <-f.Watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err := f.Watcher.Add(volume)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
