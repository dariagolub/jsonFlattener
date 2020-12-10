package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

// As sequence of keys is not guaranteed in Json, to compare exact strings, json with only 1 key is tested here
func TestProcessOneKey(t *testing.T) {
	input := "{\"a\": {\"b\": 1}}"
	expectedOutput := "{\"a.b\":1}\n"

	var stdin bytes.Buffer
	var stdout bytes.Buffer

	stdin.Write([]byte(input))

	err := Process(&stdin, &stdout)

	actualOutput := stdout.String()

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if expectedOutput != actualOutput {
		t.Errorf("Expected output: %s, actual: %s", expectedOutput, actualOutput)
	}
}

func TestProcessMoreKeys(t *testing.T) {
	input := "{\"a\": {\"b\": 1}, \"c\": \"test\"}"

	var stdin bytes.Buffer
	var stdout bytes.Buffer

	stdin.Write([]byte(input))

	err := Process(&stdin, &stdout)

	actualOutputString := stdout.String()
	var actualOutput map[string]interface{}
	json.Unmarshal([]byte(actualOutputString), &actualOutput)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if len(actualOutput) != 2 {
		t.Errorf("Actual output should contain 2 keys, but instead contains %d", len(actualOutput))
	}

	var expectedValueForKeyAB float64 =  1
	if  expectedValueForKeyAB != actualOutput["a.b"] {
		t.Errorf("Actual output should contain key 'a.b' with value 1, but instead: %v", actualOutput["a.b"])
	}
	if actualOutput["c"] != "test" {
		t.Errorf("Actual output should contain key 'c'with value 'test', but instead: %v", actualOutput["c"])
	}

}
