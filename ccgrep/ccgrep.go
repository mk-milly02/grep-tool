package ccgrep

import (
	"bytes"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	filepath string
	content  []byte
}

type Output struct {
	Path  string
	Match string
}

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

func GetFilesInDirectoryRecursively(path string) (output []string) {
	if path == "*" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path = pwd
	}

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
		if strings.Contains(line, pattern) {
			output = append(output, line)
		}
	}
	if len(output) == 0 {
		code = 1
	}
	return output, code
}

func MultipleMatch(fileInfo []file, pattern string) (result []Output, code int) {
	for _, info := range fileInfo {
		reader := bytes.NewBuffer(info.content)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			if strings.Contains(line, pattern) {
				result = append(result, Output{Path: info.filepath, Match: line})
			}
		}
	}
	if len(result) == 0 {
		code = 1
	}
	return result, code
}
