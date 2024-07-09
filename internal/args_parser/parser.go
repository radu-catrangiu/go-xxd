package argsparser

import (
	"flag"
	"log"
	"os"
	"path"
)

type Params struct {
	Infile         string
	Cols           int
	GroupSize      int
	MaxBytesToRead int
}

func ParseArgs() Params {

	groupSizePtr := flag.Int("g", 2, "Separate the output of every <n> bytes (two hex characters) by a whitespace")

	colsPtr := flag.Int("c", 16, "Format <n> octets per line")

	lenPtr := flag.Int("l", -1, "stop after <len> octets")

	flag.Parse()

	tail := flag.Args()

	infile := tail[0]

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could get current dir. Error:", err)
	}

	filePath := path.Join(currentDir, infile)

	return Params{
		Infile:         filePath,
		Cols:           *colsPtr,
		GroupSize:      *groupSizePtr,
		MaxBytesToRead: *lenPtr,
	}
}
