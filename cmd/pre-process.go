package cmd

import (
	"log"
	"os"

	"github.com/oresdev/ipcr/converter"
	"github.com/oresdev/ipcr/work"
	"github.com/spf13/cobra"
)

var (
	pre_process = &cobra.Command{
		Use:   "pre-process i o",
		Short: "Use this command with the image directory path to pre-process existing files.",
		Long:  ``,
		Run:   pre,
		Args:  cobra.ExactArgs(2),
	}
)

func pre(ccmd *cobra.Command, args []string) {
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
	// reading the catalog and picking up the task
	entries, err := os.ReadDir(i)
	if err != nil {
		log.Fatal(err)
	}
	// add an entry to the task channel
	for _, entry := range entries {
		queue <- entry.Name()
	}
	close(queue)
	// wait
	wg.Wait()
}
