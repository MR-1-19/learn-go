package main

import "fmt"

const hello = "Hello "

func Hello(name string) string {
	if name == "" {
		return hello + "world"
	}
	return hello + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
