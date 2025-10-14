package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

type contactinfo struct {
	email   string
	zipCode int
}

func main() {
	p := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactinfo{
			email:   "alex.anderson@example.com",
			zipCode: 12345,
		},
	}
	var jose person
	jose.firstName = "Silvio"
	jose.lastName = "Alvarez"
	jose.contact = contactinfo{
		email:   "silvio.alvarez@example.com",
		zipCode: 54321,
	}
	//fmt.Println(p, jose)
	//fmt.Printf("%+v\n", p)
	//fmt.Printf("%+v\n", jose)
	//p.updateName("Alexander")
	//pointerToP := &p
	//pointerToP.updateName("Alexander")
	p.updateName("Alexander")
	p.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

//func (p *person) updateName(newFirstName string) {
//	(*p).firstName = newFirstName
//}
