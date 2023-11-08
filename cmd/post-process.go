package cmd

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/oresdev/ipcr/converter"
	"github.com/oresdev/ipcr/work"
	"github.com/spf13/cobra"
)

var (
	post_process = &cobra.Command{
		Use:   "post-process i o",
		Short: "Use this command to watch your images directory and post-process the added files.",
		Long:  ``,
		Run:   post,
		Args:  cobra.ExactArgs(2),
	}
)

func post(ccmd *cobra.Command, args []string) {
	// input / output arguments
	i := args[0]
	o := args[1]
	// initializing a task pool
	queue := make(chan string)
	// initializing a worker
	worker := work.New(i, o, queue, quality, converter.New())
	// starting up our workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker.Run(i)
		}(i)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Watcher failed: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Добавляем в канал только новые события не относящиеся к измененным
				if event.Has(fsnotify.Create) && filepath.Ext(event.Name) != ".webp" {
					queue <- filepath.Base(event.Name)
				}
				log.Println("event:", filepath.Base(event.Name))

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}

	}()

	err = watcher.Add(i)
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done
}
