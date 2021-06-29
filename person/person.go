package person

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) SayHello() {
	fmt.Printf("Person's name is %s, and age is %d\n", p.name, p.age)
}

// immutable
func (p Person) Person(name string, age int) Person {
	return Person{name, age}
}

// mutable
func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}
