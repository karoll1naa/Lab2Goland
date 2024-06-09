package main

import (
	"flag"
	lab2 "github.com/karoll1naa/Lab2Goland"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "Results")
	reader          io.Reader
	writer          io.Writer
)

func main() {
	flag.Parse()

	if (*inputExpression == "") == (*inputFile == "") {
		panic("Error. Use -f or -e")
	}

	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader = file
	} else {
		reader = strings.NewReader(*inputExpression)
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}
	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}
	err := handler.Compute()
	if err != nil {
		panic(err)
	}
}
