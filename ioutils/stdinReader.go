package ioutils

import (
	"bufio"
	"os"
	"strings"
)

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
