package main

import "fmt"

const english_hello = "Hello "
const spanish_hello = "Hola "
const french_hello = "Bonjour "

func Hello(name string, language string) string {
	prefix := english_hello
	switch language {
	case "Spanish":
		prefix = spanish_hello
	case "French":
		prefix = french_hello
	default:
		prefix = english_hello
	}

	if name != "" {
		return prefix + name
	} else {
		return prefix + "world"
	}
}

func main() {
	fmt.Println("")
}
