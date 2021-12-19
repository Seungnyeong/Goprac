package main

import (
	"fmt"

	"github.com/snkim/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	word := "hello"
	definition := "Greeting"
	definition, error := dictionary.Search("first")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}

	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	hello, error := dictionary.Search(word)
	fmt.Println(hello)
}
