package main

import "./utils"

func main() {

	var inputJSON = utils.UnmarshalJSON(utils.ReadStdin())

	const keyDelimiter = "."

	var leftKey = ""
	var flattenedJSON = make(map[string]interface{})

	utils.FlattenKeyValues(inputJSON, &flattenedJSON, leftKey, keyDelimiter)

	utils.WriteJSON(flattenedJSON)

}
