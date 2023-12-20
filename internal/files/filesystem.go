package filesystem

import (
	"fmt"
	"os"
)

func Delete(filename string) {
	err := os.RemoveAll(filename)
	if err != nil {
		panic(fmt.Errorf("could not find target file to delete: %w", err))
	}
}

func DeleteAll(targets []string) {
	for _, target := range targets {
		Delete(target)
	}
}
