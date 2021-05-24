package main

import (
	"fmt"
	"encoding/json"
)

type Book struct {
	Title string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Developer bool `json:"is_developer"`
}

func main() {
	author := Author{Name: "Sean Zian", Age: 29, Developer: true}
	book := Book{Title: "Working with JSON", Author: author}

	fmt.Printf("%v\n", book)

	// byteArray, err := json.Marshal(book)
	// for prettyies print, use json.MarshalIndent() - with 2nd arg as prefix, 3rd as indent
	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}