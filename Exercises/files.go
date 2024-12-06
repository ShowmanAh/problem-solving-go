package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	err := scanDirectory("Exercises")
	if err != nil {
		log.Fatal(err)
	}
}

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				return err
			} else {
				fmt.Println(filePath)
			}
		}
	}
	return nil
}
