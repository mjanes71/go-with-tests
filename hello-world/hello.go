package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(hello("Megan", "English"))
}

func hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf(greetingPrefix(language)+"%v!", name)
}

//named return values, like this one prefix, create
//a variable in your function with that name and type
//its assigned a zero value/empty string by default
func greetingPrefix(language string) (prefix string){
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default: 
		prefix = englishHelloPrefix
	}
	return 
}
