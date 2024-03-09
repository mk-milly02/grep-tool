package main

import (
	"ccgrep/ccgrep"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()

	pattern := flag.Arg(0)
	filepath := flag.Arg(1)
	text := read_from_file(filepath)
	output := ccgrep.Match(text, pattern)

	for _, line := range output {
		fmt.Print(line)
	}
}

func read_from_file(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("%s: no such file or directory", filepath)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
