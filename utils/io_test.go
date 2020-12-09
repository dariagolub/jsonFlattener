package utils

import (
	"bytes"
	//"io/ioutil"
	"testing"
)

func TestReadUntilEmptyString(t *testing.T) {
	var stdin bytes.Buffer

	//fileData, _ := ioutil.ReadFile("resources/testData.json")
	stdin.Write([]byte("{}\n"))

	result, err := ReadUntilEmptyString(&stdin)

	if err != nil {
		t.Errorf("Error is unexpected. Should return value '{\"a\": \"test\"}'. Error: %s", err)
	}
	if result != "{\"a\": \"test\"}" {
		t.Errorf("Expected value is '{\"a\": \"test\"}', actual: %v", result)
	}

}
