package main

import (
	"./utils"
	"fmt"
	"io"
	"log"
	"os"
)

const keyDelimiter = "."

func main() {
	err := Process(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

}

// Process takes input json data and writes flattened json to output
func Process(input io.Reader, output io.Writer) error {
	rawData, err := utils.ReadUntilEOF(input)
	if err != nil {
		return err
	}

	inputJSON, err := utils.UnmarshalJSON(rawData)
	if err != nil {
		return err
	}

	var leftKey = ""
	var flattenedJSON = make(map[string]interface{})

	err = utils.FlattenKeyValues(inputJSON, &flattenedJSON, leftKey, keyDelimiter)
	if err != nil {
		return err
	}

	result, err := utils.MarshalJSON(flattenedJSON)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(output, "%s\n", result)
	if err != nil {
		return err
	}

	return nil
}
