package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//WITHOUT INTERFACE WOULD CAUSE AN ERROR AS
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sp spanishBot) {
// 	fmt.Println(sp.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY CUSTOM LOGIC FOR GENERATING AN ENGLISH GREETING
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// VERY CUSTOM LOGIC FOR GENERATING AN SPANISH GREETING
	return "Hola!"
}
