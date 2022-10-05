package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main(){
	fmt.Println(Hello("Megan"))
}

func Hello(name string) string{
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf(englishHelloPrefix + "%v!", name)
}