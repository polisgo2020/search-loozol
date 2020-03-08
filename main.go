package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

type indexReverse map[string][]string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func indexingFolder(filePath string) {
	files, err := ioutil.ReadDir(filePath)
	check(err)
	index := make(indexReverse)

	for _, file := range files {
		doc, err := ioutil.ReadFile(filePath + string(os.PathSeparator) + file.Name())
		check(err)

		line := strings.Fields(string(doc))

		for _, word := range line {
			word = strings.TrimFunc(string(word), func(f rune) bool {
				return !unicode.IsLetter(f) && !unicode.IsNumber(f)
			})
			index[word] = append(index[word], file.Name())
		}
	}
	_, err = os.Create("output.txt")
	check(err)
	output, err := json.Marshal(index)
	check(err)
	if err := ioutil.WriteFile("output.txt", output, 0644); err != nil {
		panic(err)
	}
}

func main() {
	var filePath = "filePath"
	indexingFolder(filePath)
}
