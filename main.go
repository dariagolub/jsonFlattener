package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var stringBuilder strings.Builder

	for {
		line, _ := reader.ReadString('\n')

		if len(line) == 0 {
			break
		}

		stringBuilder.WriteString(line)

	}

	rawData := []byte(stringBuilder.String())

	var inputJson map[string]interface{}
	err := json.Unmarshal(rawData, &inputJson)
	
	if err != nil {
		panic(err) //todo
	}

	var leftKey = ""
	keyDelimiter := "."
	var flattenedJson = make(map[string]interface{})

	flattenJson(inputJson, &flattenedJson, leftKey, keyDelimiter)
	for key, value := range flattenedJson {
		fmt.Printf("%v:%v\n", key, value)
	}

}

func flattenJson(inputJSON map[string]interface{}, flattened *map[string]interface{}, leftKey string, keyDelimiter string) {
	for key, value := range inputJSON {
		key := leftKey + key

		if _, ok := value.(map[string]interface{}); ok {
			flattenJson(value.(map[string]interface{}), flattened, key+keyDelimiter, keyDelimiter)
		} else {
			(*flattened)[key] = value
		}
	}
}
