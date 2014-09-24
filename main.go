package main

import (
	"fmt"
	"flag"
)

func main() {
	var directory string

	flag.StringVar(&directory, "directory", ".", "Root directory for search")

	flag.Parse()

	fmt.Println(directory)
}
