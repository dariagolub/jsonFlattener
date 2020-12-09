package main

import (
	"bytes"
	"testing"
)

func TestProcess(t *testing.T) {
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
