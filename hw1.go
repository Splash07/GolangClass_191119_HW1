package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Student struct
	type Student struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
		ClassName string `json:"class_name"`
	}
	var serializedStruct []Student

	jsonPayload := []byte(`[
							{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
							{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
							]`)

	err := json.Unmarshal(jsonPayload, &serializedStruct)

	if err != nil {
		panic(err)
	}

	for _, o := range serializedStruct {
		fmt.Println(o)
	}
}
