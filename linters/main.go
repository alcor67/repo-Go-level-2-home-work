package main

import (
	"encoding/json"
	"fmt"
)

// MyStruct is struct{}
type MyStruct struct{}

// MarshalJSON is (f MyStruct) MarshalJSON() ([]byte, error)
func (f MyStruct) MarshalJSON() ([]byte, error) {
	return []byte(`{"a": 1}`), nil
}
func main() {
	j, _ := json.Marshal(MyStruct{})
	fmt.Printf("type: %T\nvalue: %+v\n", j, j)
	fmt.Println(string(j))
}
