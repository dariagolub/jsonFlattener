package main

import "testing"

func TestUnmarshalJsonCorrect(t *testing.T) {
	rawData := "{\"a\": \"b\"}";
	jsonData := UnmarshalJson(rawData)

	if jsonData["a"] != "b" {
		t.Error("Json should contain key \"a\"")
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

	FlattenJson(jsonData, &flattenedJson, "", ".")

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
