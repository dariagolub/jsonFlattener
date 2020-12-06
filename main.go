package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter json:")

	var stringBuilder strings.Builder

	for {
		line, _ := reader.ReadString('\n')

		if len(line) ==0 {
			break
		}

		stringBuilder.WriteString(line)

	}

	fmt.Printf(stringBuilder.String())
}
