package main

import (
	"fmt"
    "errors"
	"github.com/bwyltcaxe/golang_employees/internal/person"
)

func main() {
	personKey := person.Person{
		Name: "John Sick",
		Gender: "man",
		ID: 0x8084,
	}

	err := personKey.SetName("Jane Wick")
	if err != nil {
		fmt.Println("Error setting name:", err)
		return
	}

	err = personKey.SetGender("woman")
	if err != nil {
		fmt.Println("Error setting gender:", err)
		return
	}

	err = personKey.SetID(0x8090)
	if err != nil {
		fmt.Println("Error setting ID:", err)
		return
	}

	workersSalary := make(map[person.Person]int)
	workersSalary[personKey] = 100

	person1 := person.Person{
		Name: "Bob Marley",
		Gender: "man",
		ID: 0x8080,
	}
	workersSalary[person1] = 200

	person2 := person.Person{
		Name: "Robert Paulson",
		Gender: "man",
		ID: 0x7070,
	}
	workersSalary[person2] = 300

	fmt.Println("=========================")
	fmt.Println("Salaries: ", workersSalary)

	delete(workersSalary, personKey)

	fmt.Println("=========================")
	fmt.Println("New salaries: ", workersSalary)

	fmt.Println("=========================")
	var personsArr [2]person.Person
	person3 := person.Person{
		Name: "Ivan Ivanov",
		Gender: "man",
		ID: 0x7070,
	}
	personsArr[0] = personKey
	personsArr[1] = person3

	for _, v := range personsArr {
		fmt.Println(">> [", v.Name, "]", ": ", v.ID)
	}

	// Let's use errors.Is
	err = personKey.SetGender("unknown")
	if err != nil {
		if errors.Is(err, person.ErrInvalidGender) {
			fmt.Println("Invalid gender:", err)
		} else {
			fmt.Println("Error setting gender:", err)
		}
	}
}

