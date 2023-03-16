package main

import (
	"flag"
	"fmt"
	"html-parser/link"
	"os"
	"strings"
)

const defaultFilePath = "./html_samples/sample1.html"

func main() {
	var filePath string
	flag.StringVar(&filePath, "path", defaultFilePath, "The filepath to the HTML that will be parsed")
	flag.Parse()

	data, err := readHtmlFromFile(filePath)
	if err != nil {
		panic(err)
	}
	r := strings.NewReader(data)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

func readHtmlFromFile(filePath string) (string, error) {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}