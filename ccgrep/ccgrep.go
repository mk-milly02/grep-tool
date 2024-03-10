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
		if strings.Contains(line, pattern) {
			output = append(output, line)
		}
	}
	if len(output) == 0 {
		code = 1
	}
	return output, code
}
