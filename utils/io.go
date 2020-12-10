package utils

import (
	"bufio"
	"io"
	"strings"
)

// ReadUntilEOF reads input from the io.Reader
func ReadUntilEOF(buffer io.Reader) (string, error) {
	reader := bufio.NewReader(buffer)

	var stringBuilder strings.Builder

	for {
		line, err := reader.ReadString('\n')

		stringBuilder.WriteString(line)

		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

	}

	rawData := stringBuilder.String()

	return rawData, nil
}