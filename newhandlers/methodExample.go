package newhandlers

import "fmt"

type Person struct {
	name     string
	phone    int
	location string
}

func (p *Person) SetPer() {
	fmt.Scan(&p.name)
	fmt.Scan(&p.phone)
	fmt.Scan(&p.location)
}

// func MyPerson() Person {
// 	per := Person{"siddu", 973117, "bangalore"}
// 	return per
// }

func (p *Person) GetPer() {
	fmt.Println("Name of the Person %v", p.name)
	fmt.Println("Phone number of the person %v", p.phone)

}
