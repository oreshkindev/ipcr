package main

import (
	"fmt"

	"github.com/oresdev/ipcr/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
