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
	i := flag.Bool("i", false, "allows case insensitive search")
	flag.Parse()

	pattern := flag.Arg(0)
	path := flag.Arg(1)

	var text []byte
	var code int
	var result []string

	switch path {
	case "":
		text = ccgrep.ReadFromStdIn()
		switch {
		case *v:
			result, code = ccgrep.MatchInversely(text, pattern)
		case *i:
			result, code = ccgrep.MatchCaseInsensitive(text, pattern)
		default:
			result, code = ccgrep.Match(text, pattern)
		}
		for _, line := range result {
			fmt.Print(line)
		}
		os.Exit(code)

	default:
		if filepath.Ext(path) != "" { // is a file
			text = ccgrep.ReadFromFile(path)
			switch {
			case *v:
				result, code = ccgrep.MatchInversely(text, pattern)
			case *i:
				result, code = ccgrep.MatchCaseInsensitive(text, pattern)
			default:
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
