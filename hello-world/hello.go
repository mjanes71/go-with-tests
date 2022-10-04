package main

import "fmt"

func main(){
	fmt.Println(Hello("Megan"))
}

func Hello(name string) string{
	return fmt.Sprintf("Hello, %v!", name)
}