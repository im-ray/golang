package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

// struct inside struct embedding struct
type contactInfo struct {
	email string
	zipCode int
}

type person2 struct {
	firstName string
	lastName string
	contact contactInfo
}
// a simple toungetwister even if we write only contactInfo for person2 struct 
// its equivalent to {contactInfo contactInfo}


func main() {

	// two different ways to define structs
	alex := person{"Alex", "Anderson"}
	// by default go assumes position variable value in order
	fmt.Println(alex.firstName)
	alex = person{firstName: "Alex", lastName: "Anderson"}
	// better way to define structs
	fmt.Println(alex)


	var alex_1 person
	// default value is 0 value depends on type
	// string --> "",   int--> 0, float--> 0,   bool--> false
	fmt.Println(alex_1) // {  }
	fmt.Printf("%+v", alex_1) // {firstName: lastName:}


	// using the embedded struct
	jim := person2{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@gmail.com", 
			zipCode: 94000,
		},
	}// every propery follow ,
	fmt.Printf("%+v", jim)
	jim.print()
	jim.updateName("jimmy")
	jim.print()

	jimpointer := &jim
	jimpointer.updateName_pointer("Jimmy")


}

// golang is pass by value
func (p person2) updateName(newFirstName string) {
	p.firstName = newFirstName
	// ❌ This modifies a copy, not the original struct --> pass by value 
}
// *type_ vs *variable of type_
func (pointerToperson2 *person2) updateName_pointer(newFirstName string) {
	(*pointerToperson2).firstName = newFirstName 
}


// write a receiver function for struct type
func (p person2) print() {
	fmt.Printf("%+v", p)
}

// pointer shortcut
// go allows jim.updateName_pointer("jimmy") as well struct_type can call 
// to pointer of struct_type


// gotchas with pointers
// https://dev.to/dawkaka/go-arrays-and-slices-a-deep-dive-dp8
func main_1() {
	mySlice := []string{"Hi", "There", "How", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}
func updateSlice(s []string) {
	s[0] = "Bye"
}
// Value Types: {int, float, string, bool, structs} --> need to use pointers
// Reference Types: {slices, maps, channels, pointers, functions}







