package ccgrep

import (
	"bytes"
	"strings"
)

func Match(text []byte, pattern string) []string {
	reader := bytes.NewBuffer(text)
	var output []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(line, pattern) {
			output = append(output, line)
		}
	}
	return output
}
