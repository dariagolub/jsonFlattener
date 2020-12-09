package utils

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON converts string JSON to map
func UnmarshalJSON(rawData string) (map[string]interface{}, error) {
	var inputJSON map[string]interface{}
	err := json.Unmarshal([]byte(rawData), &inputJSON)

	if err != nil {
		return nil, fmt.Errorf("json input is not correct. Error: %s", err)
	}

	return inputJSON, nil
}

// MarshalJSON converts map to JSON string
func MarshalJSON(jsonData map[string]interface{}) (string, error) {
	json, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

// FlattenKeyValues to flatten key-values pairs
// It's a recursive function to iterate through keys and flatten key-value pairs if value represents a nested structure
func FlattenKeyValues(inputJSON map[string]interface{}, flattened *map[string]interface{}, leftKey string, keyDelimiter string) error {
	for key, value := range inputJSON {
		key := leftKey + key

		if ok := isTypeArray(value); ok {
			return fmt.Errorf("flattening json with array is not supported. Wrong key: %v with value: %v", key, value)
		} else if ok := isTypeInnerJSON(value); ok {
			err := FlattenKeyValues(value.(map[string]interface{}), flattened, key+keyDelimiter, keyDelimiter)
			if err != nil {
				return err
			}
		} else {
			(*flattened)[key] = value
		}
	}

	return nil
}

func isTypeArray(value interface{}) bool {
	_, ok := value.([]interface{})
	return ok
}

func isTypeInnerJSON(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}
