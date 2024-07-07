package main

import (
	"log"
	"os"

	argsparser "github.com/radu-catrangiu/go-xxd/internal/args_parser"
	reader "github.com/radu-catrangiu/go-xxd/internal/reader"
)

func main() {
	params := argsparser.ParseArgs()

	file, err := os.Open(params.Infile)
	if err != nil {
		log.Fatalln("Could not open the file. Error:", err)
	}

	reader.Read(*file, params.Cols, params.GroupSize)

	err = file.Close()
	if err != nil {
		log.Fatalln("Could not close file. Error:", err)
	}
}
