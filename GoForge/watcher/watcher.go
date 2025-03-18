package watcher

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/radovskyb/watcher"
)

func StartWatcher(dirs ...string) {
	w := watcher.New()
	w.SetMaxEvents(1)

	for _, dir := range dirs {
		if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			w.Add(path)
			return nil
		}); err != nil {
			log.Printf("Watcher error: %v", err)
		}
	}

	go func() {
		for {
			select {
			case <-w.Event:
				log.Println("File changed, restarting server...")
				cmd := exec.Command("pkill", "-f", "advgofront")
				cmd.Run()
				go func() {
					exec.Command("go", "run", "main.go").Start()
				}()
			case err := <-w.Error:
				log.Printf("Watcher error: %v", err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Printf("Watcher failed: %v", err)
	}
}
