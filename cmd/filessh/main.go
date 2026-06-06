package main

import (
	"fmt"
	"os"

	"github.com/LSN0WM4N/filessh/pkg/explorer"
)

func main() {
	var entries []explorer.Entry
	var actualDir string

	if wd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		actualDir = wd
		entries, err = explorer.ReadDir(wd)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Actual dir: %s \n", actualDir)
	for _, entry := range entries {
		icon := "*"
		if entry.IsDir {
			icon = "/"
		}
		fmt.Printf("[%s] %s\n", icon, entry.Name)
	}
}
