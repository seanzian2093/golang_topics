package main

import (
	"fmt"
	"encoding/json"
)

type SencorReading struct {
	Name string `json:"name"`
	Capacity int `json:"capacity"`
	Time string `json:"time"`
}

func main() {
	// backtick for raw string literal, i.e. no special character, e.g. \n will be \ and n, not a new line
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2021-05-20T19:07:282"}`
	
	var reading SencorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", reading)


}