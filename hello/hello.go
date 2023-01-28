package main

import "fmt"

const englishGreetingPrefix = "Hello, "
const spanishGreetingPrefix = "Hola, "
const frenchGreetingPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) (prefix string) {

	switch language {
	case "Spanish":
		prefix = spanishGreetingPrefix
	case "French":
		prefix = frenchGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
