package main

import (
	"fmt"

	"github.com/meowailand/MeowCoin/person"
)

type innerPerson struct {
	name string
	age  int
}

func (ip innerPerson) sayHello() {
	fmt.Printf("My name is %s, and my age is %d\n", ip.name, ip.age)
}

func main() {

	var p1 innerPerson = innerPerson{"meow", 14}
	p1.sayHello()

	p2 := person.Person{}.Person("meow1", 15)
	p2.SayHello()

	p3 := person.Person{}
	p3.SetPerson("meow2", 16)
	p3.SayHello()
}
