package main

import (
	"./ioutils"
	"encoding/json"
	"log"
)

func main() {

	var inputJson = UnmarshalJson(ioutils.ReadStdin())

	const keyDelimiter = "."

	var leftKey = ""
	var flattenedJson = make(map[string]interface{})

	FlattenJson(inputJson, &flattenedJson, leftKey, keyDelimiter)

	ioutils.WriteJson(flattenedJson)

}

func UnmarshalJson(rawData string) map[string]interface{} {
	var inputJson map[string]interface{}
	err := json.Unmarshal([]byte(rawData), &inputJson)

	if err != nil {
		log.Fatalf("Json input is not correct. Error: %s" , err) //todo
	}

	return inputJson
}

func FlattenJson(inputJSON map[string]interface{}, flattened *map[string]interface{}, leftKey string, keyDelimiter string) {
	for key, value := range inputJSON {
		key := leftKey + key

		if ok := isTypeArray(value); ok {
			log.Fatalf("Flattening json with array is not supported. Wrong key: %v with value: %v", key, value) //todo
		} else if ok := isTypeInnerJson(value); ok {
			FlattenJson(value.(map[string]interface{}), flattened, key+keyDelimiter, keyDelimiter)
		} else {
			(*flattened)[key] = value
		}
	}
}

func isTypeArray(value interface{}) bool {
	_, ok := value.([]interface{})
	return ok
}

func isTypeInnerJson(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}
