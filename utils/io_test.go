package utils

import (
	"bytes"
	"testing"
)

func TestReadUntilEmptyString(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("{\"a\": \"test\"}"))

	actual, err := ReadUntilEOF(&stdin)

	if err != nil {
		t.Errorf("Error is unexpected. Should return value '{\"a\": \"test\"}'. Error: %s", err)
	}

	expected := "{\"a\": \"test\"}"

	if expected != actual {
		t.Errorf("Expected value is %s, actual: %s", expected, actual)
	}

}
