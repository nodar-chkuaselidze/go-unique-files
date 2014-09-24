package main

import (
	"fmt"
	"flag"
	"os"
	"path/filepath"
	"crypto/sha1"
	"io/ioutil"
	"encoding/hex"
)

const defaultDirectory = "."

var result = make(map[string][]string)

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

	var files = make(map[string][]string)

	err := filepath.Walk(directory, func (path string, fileInfo os.FileInfo, err error) error {
		str, werr := walkFn(path, fileInfo, err)

		if werr != nil || len(str) == 0 {
			return werr
		}

		if _, ok := files[str]; ok {
			files[str] = append(files[str], path)
		} else {
			files[str] = append(make([]string, 0, 10), path)
		}

		return werr
	});

	for hash, filePaths := range files {
		fmt.Println(hash)
		for _, file := range filePaths {
			fmt.Println("\t", file)
		}
	}

	if err != nil {
		panic(err)
	}
}

func walkFn(path string, fileInfo os.FileInfo, err error) (string, error) {
	if err != nil || fileInfo.IsDir() {
		return "", err
	}

	bs, err := getFileSha1(path, fileInfo)

	if err != nil {
		return "", err
	}

	str := hex.EncodeToString(bs)

	return str, nil
}

func getFileSha1(path string, fileInfo os.FileInfo) ([]byte, error) {
	bs, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	h := sha1.New()
	h.Write(bs)
	return h.Sum(nil), nil
}
