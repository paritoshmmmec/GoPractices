package main

import (
	"fmt"
	"github.com/sample/Logger"
	"github.com/sample/Model"
)

func main() {
	goku := initialize("Player 1", 0)

	goku.Introduce("testing")

	Logger.Log("Initializing first player")

	fmt.Printf("My Name is %s\n", goku.Person.Name)
	fmt.Printf("My Power is %d\n", goku.Power)

	goku.Person.Name = "Paritosh"
	goku.Power = 9000

	fmt.Printf("My Name is %s\n", goku.Contact.Name)
	fmt.Printf("My Power is %d\n", goku.Power)

	Logger.Log("Recharging the new power")
	goku.Person.Introduce("Calling through mixing")
	Super(&goku)

	fmt.Printf("My Name is %s\n", goku.Contact.Name)
	fmt.Printf("My Power is %d\n", goku.Power)
}

func Super(s *Model.Saiyan) {
	s.Power += 10000
}

func initialize(name string, power int) Model.Saiyan {
	return Model.Saiyan{
		Contact: &Model.Contact{
			Name: name},
		Person: &Model.Person{
			Name: name},
		Power: power,
	}
}
