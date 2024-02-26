package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
)

func Hello(name string, dialect string) string {
	if name == "" {
		name = "World"
	}
	return getGreeting(name, dialect)
}

func getGreeting(name string, dialect string) string {
	switch dialect {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Intuji", ""))
}
