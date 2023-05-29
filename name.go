// package main

// import (
// 	"fmt"
// )

// type contactInfo struct {
// 	pincode int
// 	email   string
// }

// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }

// func main() {
// 	kausha := person{
// 		firstName: "Kaushha",
// 		lastName:  "Trivedi",
// 		contactInfo: contactInfo{
// 			pincode: 401202,
// 			email:   "lala@gmail.com",
// 		},
// 	}

// 	// kaushaPointer := &kausha
// 	// kaushaPointer.updateName("Kausha")

// 	kausha.updateName("Kausha")

// 	// fmt.Printf("%+v", kausha)
// 	kausha.print()
// 	// Nency := person{"Nency", "Batada"}
// 	// Reshma := person{firstName: "Reshma", lastName: "Batada"}
// 	// fmt.Println(Nency, Reshma)
// 	// testing()
// }

// func testing() {
// 	var dhvani person

// 	dhvani.firstName = "dhvani"
// 	dhvani.lastName = "shah"

// 	fmt.Println(dhvani)
// 	fmt.Printf("%+v", dhvani)
// }

// func (pointertoperson *person) updateName(newFirstName string) {
// 	(*pointertoperson).firstName = newFirstName
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }
