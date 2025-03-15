package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
	
	// printGreeting(sb)
	//cannot use sb (variable of type spanishBot) as englishBot value in argument to printGreetingcompilerIncompatibleAssign


}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}


// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }


//printGreeting redeclared in this blockcompilerDuplicateDecl
//main.go .....: other declaration of printGreeting

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}