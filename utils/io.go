package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadStdin reads input from stdin
func ReadStdin() string {
	reader := bufio.NewReader(os.Stdin)

	var stringBuilder strings.Builder

	for {
		line, _ := reader.ReadString('\n')

		if len(line) == 0 {
			break
		}

		stringBuilder.WriteString(line)

	}

	rawData := stringBuilder.String()

	return rawData
}

// WriteJSON to write map with JSON key-values to standard output
func WriteJSON(flattenedJSON map[string]interface{}) {
	for key, value := range flattenedJSON {
		fmt.Printf("%v:%v\n", key, value)
	}
}
