package cmd

import (
	"sync"

	"github.com/spf13/cobra"
)

const def = 4

var (
	quality int
	wg      sync.WaitGroup
	workers int

	commands = &cobra.Command{
		Use:           "ipcr",
		Short:         "ipcr – command-line tool for Image-Processor",
		Long:          `Image-Processor is a CLI microservice that extends the capabilities of image processing. Use this tool to automatically decode / compress PNG and JPEG files and then convert them to WEBP and JPG formats.`,
		Version:       "1.0.0",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return commands.Execute()
}

func init() {
	// determine the number of wokers if you passed the argument
	commands.PersistentFlags().IntVarP(&workers, "workers", "w", def, "determine the number of wokers")
	// determine the quality of images when compressed
	commands.PersistentFlags().IntVarP(&quality, "quality", "q", 60, "compressed image quality")
	commands.AddCommand(post_process)
	commands.AddCommand(pre_process)
}
