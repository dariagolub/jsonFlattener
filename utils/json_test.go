package utils

import "testing"

func TestUnmarshalJsonCorrect(t *testing.T) {
	rawData := "{\"a\": \"b\"}";
	jsonData, err := UnmarshalJSON(rawData)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if jsonData["a"] != "b" {
		t.Error("Json should contain key \"a\"")
	}
}

func TestUnmarshalJsonError(t *testing.T) {
	rawData := "{\"not a json\"}";
	jsonData, err := UnmarshalJSON(rawData)

	if err == nil {
		t.Error("Error is expected")
	}
	if jsonData != nil {
		t.Error("Nil result is expected")
	}
}

func TestFlattenJsonCorrect(t *testing.T) {
	jsonData := map[string]interface{}{
		"a": 1,
		"c": map[string]interface{}{
			"d": 3,
			"e": map[string]interface{}{
				"g": 2.52,
			},
		},
	}

	var flattenedJson = make(map[string]interface{})

	err := FlattenKeyValues(jsonData, &flattenedJson, "", ".")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(flattenedJson) != 3 {
		t.Errorf("Amount of keys in flattened json should be 3, but actual is %d", len(flattenedJson))
	}
	if flattenedJson["a"] != 1 {
		t.Errorf("Value of key 'a' should be 1, but actual is %v", flattenedJson["a"])
	}
	if flattenedJson["c.d"] != 3 {
		t.Errorf("Value of key 'c.d' should be 3, but actual is %v", flattenedJson["c.d"])
	}
	if flattenedJson["c.e.g"] != 2.52 {
		t.Errorf("Value of key 'c.e.g' should be 2.52, but actual is %v", flattenedJson["c.e.g"])
	}

}

func TestFlattenJsonArrayNotSupportedError(t *testing.T) {
	jsonData := map[string]interface{}{
		"a": 1,
		"c": [2]int{1, 2},
	}

	var flattenedJson = make(map[string]interface{})

	err := FlattenKeyValues(jsonData, &flattenedJson, "", ".")

	if err != nil {
		t.Errorf("Error is expected")
	}
}
