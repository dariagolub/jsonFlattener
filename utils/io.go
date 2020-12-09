package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ReadUntilEmptyString reads input from stdin
func ReadUntilEmptyString(buffer io.Reader) (string, error) {
	reader := bufio.NewReader(buffer)

	var stringBuilder strings.Builder

	for {
		line, err := reader.ReadString('\n')

		stringBuilder.WriteString(line)

		if err == io.EOF {
			break
		}

	}

	rawData := stringBuilder.String()

	return rawData, nil
}

// WriteJSON to write map with JSON key-values to standard output
func WriteJSON(flattenedJSON map[string]interface{}) {
	for key, value := range flattenedJSON {
		fmt.Printf("%v:%v\n", key, value)
	}
}
