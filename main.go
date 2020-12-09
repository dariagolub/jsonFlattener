package main

import (
	"./utils"
	"log"
	"os"
)

func main() {

	rawData, err := utils.ReadUntilEmptyString(os.Stdin)
	if err != nil {
		log.Fatalf("Json input is not correct. Error: %s" , err)
	}
	//log.Fatalf("Json input is not correct. Error: %s" , err) //todo

	inputJSON, _ := utils.UnmarshalJSON(rawData)

	const keyDelimiter = "."

	var leftKey = ""
	var flattenedJSON = make(map[string]interface{})

	_ = utils.FlattenKeyValues(inputJSON, &flattenedJSON, leftKey, keyDelimiter)

	utils.WriteJSON(flattenedJSON)

}
