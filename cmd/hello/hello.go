package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const teReo = "Te Reo"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const teReoHelloPrefix = "Kia ora, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case teReo:
		prefix = teReoHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// this is an example of a naked return: returns the named return values
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
