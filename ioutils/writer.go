package ioutils

import "fmt"

func WriteJson(flattenedJson map[string]interface{}) {
	for key, value := range flattenedJson {
		fmt.Printf("%v:%v\n", key, value)
	}
}
