package utils

import (
	"encoding/json"
	"log"
)

// UnmarshalJSON converts string JSON to map
func UnmarshalJSON(rawData string) map[string]interface{} {
	var inputJSON map[string]interface{}
	err := json.Unmarshal([]byte(rawData), &inputJSON)

	if err != nil {
		log.Fatalf("Json input is not correct. Error: %s" , err) //todo
	}

	return inputJSON
}

// FlattenKeyValues to flatten key-values pairs
// It's a recursive function to iterate through keys and flatten key-value pairs if value represents a nested structure
func FlattenKeyValues(inputJSON map[string]interface{}, flattened *map[string]interface{}, leftKey string, keyDelimiter string) {
	for key, value := range inputJSON {
		key := leftKey + key

		if ok := isTypeArray(value); ok {
			log.Fatalf("Flattening json with array is not supported. Wrong key: %v with value: %v", key, value) //todo
		} else if ok := isTypeInnerJSON(value); ok {
			FlattenKeyValues(value.(map[string]interface{}), flattened, key+keyDelimiter, keyDelimiter)
		} else {
			(*flattened)[key] = value
		}
	}
}

func isTypeArray(value interface{}) bool {
	_, ok := value.([]interface{})
	return ok
}

func isTypeInnerJSON(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}

