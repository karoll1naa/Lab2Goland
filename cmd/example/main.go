package main

import (
	"flag"
	"fmt"
	lab2 "github.com/karoll1naa/Lab2Goland"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	res, _ := lab2.PostfixToInfix("+ 2 2")
	fmt.Println(res)
}
