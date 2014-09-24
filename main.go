package main

import (
	"fmt"
	"flag"
	"os"
	"path/filepath"
)

const defaultDirectory = "."

func main() {
	flag.Parse()
	directory := flag.Arg(0)

	if len(directory) == 0 {
		directory = defaultDirectory
	}

	searchFiles(directory)
}

func searchFiles(directory string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	err := filepath.Walk(directory, walkFn);

	if err != nil {
		panic(err)
	}
}

func walkFn(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(path)

	return nil
}
