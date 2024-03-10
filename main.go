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
	v := flag.Bool("v", false, "inverts the search excluding and result that matches")
	flag.Parse()

	pattern := flag.Arg(0)
	path := flag.Arg(1)

	var text []byte
	var code int
	var result []string

	switch path {
	case "":
		text = ccgrep.ReadFromStdIn()
		if *v {
			result, code = ccgrep.MatchInversely(text, pattern)
		} else {
			result, code = ccgrep.Match(text, pattern)
		}
		for _, line := range result {
			fmt.Print(line)
		}
		os.Exit(code)

	default:
		if filepath.Ext(path) != "" { // is a file
			text = ccgrep.ReadFromFile(path)
			if *v {
				result, code = ccgrep.MatchInversely(text, pattern)
			} else {
				result, code = ccgrep.Match(text, pattern)
			}
			for _, line := range result {
				fmt.Print(line)
			}
			os.Exit(code)
		}

		//is a folder and -r is specified
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
}
