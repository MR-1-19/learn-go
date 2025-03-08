package main

import (
	"fmt"
)

const english_hello = "Hello "
const spanish_hello = "Hola "
const french_hello = "Bonjour "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return choose_language(language) + name
}

func choose_language(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanish_hello
	case "French":
		prefix = french_hello
	default:
		prefix = english_hello
	}
	return
}

func main() {
	fmt.Println("")
}
