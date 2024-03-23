package ccgrep

import (
	"bytes"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func ReadFromFile(filepath string) []byte {
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

func ReadFromStdIn() []byte {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func GetFilesInDirectoryRecursively(path string) (output []string, parentDir string) {
	if path == "*" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path = pwd
	}
	parentDir = path
	err := filepath.WalkDir(path,
		func(path string, d fs.DirEntry, err error) error {
			if strings.Contains(path, "txt") {
				output = append(output, path)
			}
			return nil
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Match(text []byte, pattern string) (output []string, code int) {
	reader := bytes.NewBuffer(text)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		switch pattern {
		case "\\d":
			if strings.ContainsAny(line, "1234567890") {
				output = append(output, line)
			}
		case "\\w":
			has_symbol := 0
			for _, char := range line {
				if unicode.IsSymbol(char) || unicode.IsPunct(char) {
					has_symbol = 1
					break
				}
			}
			if has_symbol == 0 {
				output = append(output, line)
			}

		default:
			if strings.Contains(line, pattern) {
				output = append(output, line)
			}
		}
	}
	if len(output) == 0 {
		code = 1
	}
	return output, code
}

func MatchInversely(text []byte, pattern string) (output []string, code int) {
	reader := bytes.NewBuffer(text)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if !strings.Contains(line, pattern) {
			output = append(output, line)
		}
	}
	if len(output) == 0 {
		code = 1
	}
	return output, code
}
