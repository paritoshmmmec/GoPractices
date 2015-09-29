package Model

import (
	"fmt"
)

type Saiyan struct {
	*Contact
	*Person
	Power int
}

type Person struct {
	Name string
}

type Contact struct {
	Name string
}

func (p *Person) Introduce(message string) {
	fmt.Printf("Hi, I'm %s %s \n", p.Name, message)
}

func (p *Saiyan) Introduce(message string) {
	fmt.Printf("Hi, I'm Saiyan %s %s \n", p.Person.Name, message)
}
