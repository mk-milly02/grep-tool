package main

import (
	"ccgrep/ccgrep"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	r := flag.Bool("r", false, "to recurse a directory tree")
	flag.Parse()

	pattern := flag.Arg(0)
	path := flag.Arg(1)

	var text []byte
	var code int
	var result []string

	if path != "*" && filepath.Ext(path) != "" { // is a file
		text = ccgrep.ReadFromFile(path)
		result, code = ccgrep.Match(text, pattern)
		for _, line := range result {
			fmt.Print(line)
		}
		os.Exit(code)
	}

	if *r {
		fs := ccgrep.GetFilesInDirectoryRecursively(path)
		for _, f := range fs {
			text = ccgrep.ReadFromFile(f)
			result, code = ccgrep.Match(text, pattern)
			for _, line := range result {
				fmt.Printf("%s:%s", filepath.Base(f), line)
			}
		}
		os.Exit(code)
	}
}
