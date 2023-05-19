package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func main() {
	johnson := person{
		firstName: "Johnson",
		lastName:  "Allison",
		contactInfo: contactInfo{
			email:   "Johnson.Allison@gmail.com",
			zipCode: 90005,
		},
	}

	johnson.print()
	johnson.updateName("Bruno")
	johnson.print()
}

// Pointers: the pointer is *pointer
// Please note any value with struct, bool, int, string, float when passed around in a func
// is being passed as a value, there4 Go will create a new copy for it and if there
// is any update to the variable in the func, it will not change the original value.
// But if it is map, slice, array, there will be passed around by reference, so any change
// to it in a func, it will reflect the changes in the original value, see below for the
// implementation of a pointer when passing value type into a func

func (p *person) updateName(newName string) {
	(p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println() // for new line
}
