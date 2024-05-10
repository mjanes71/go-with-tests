package main

import "fmt"

// common to group constants and leave a line between related groups
const (

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("Megan", "English"))
}

func Hello(name, language string) string {
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
