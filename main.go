package main

import (
	"ccgrep/ccgrep"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	pattern := flag.Arg(0)
	filepath := flag.Arg(1)
	text := ccgrep.ReadFromFile(filepath)
	output, code := ccgrep.Match(text, pattern)

	for _, line := range output {
		fmt.Print(line)
	}
	os.Exit(code)
}
