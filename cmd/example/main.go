package main

import (
	"flag"
	"os"
	"strings"
	lab2 "github.com/Rabotiagi/SA_Lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputDirectory = flag.String("f", "", "File path to the input file")
	outputDirectory = flag.String("o", "", "Path to the file to place the output in")
)

func main() {
	flag.Parse()

	if *inputExpression != "" && *inputDirectory != "" {
		panic("only one flag can be used for input specifying")
	}

	if *inputExpression == "" && *inputDirectory == "" {
		panic("no input value was specified")
	}

	handler := &lab2.ComputeHandler {
		Input: strings.NewReader(*inputExpression),
		Output: os.Stdout,
	}

	if *inputDirectory != "" {
		handlerInput, err1 := os.OpenFile(*inputDirectory, 0, 0777)
		if (err1 != nil) {
			panic(err1)
		}
		handler.Input = handlerInput
	}

	if *outputDirectory != "" {
		handlerOutput, err2 := os.OpenFile(*outputDirectory, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		if (err2 != nil) {
			panic(err2)
		}
		handler.Output = handlerOutput
	}

	err := handler.Compute()
	if (err != nil) {
		panic(err)
	}
}