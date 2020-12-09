package utils

import "testing"

func TestUnmarshalJsonCorrect(t *testing.T) {
	rawData := "{\"a\": \"b\"}";
	jsonData := UnmarshalJSON(rawData)

	if jsonData["a"] != "b" {
		t.Error("Json should contain key \"a\"")
	}

}
