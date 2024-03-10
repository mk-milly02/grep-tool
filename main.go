package main

import (
	"ccgrep/ccgrep"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		fs, pd := ccgrep.GetFilesInDirectoryRecursively(path)
		for _, f := range fs {
			text = ccgrep.ReadFromFile(f)
			result, code = ccgrep.Match(text, pattern)
			for _, line := range result {
				px, _ := strings.CutPrefix(f, pd)
				fmt.Printf("%s:%s", strings.TrimPrefix(px, "\\"), line)
			}
		}
		os.Exit(code)
	}
}
