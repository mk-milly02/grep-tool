package ccgrep

import (
	"bytes"
	"io"
	"log"
	"os"
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
