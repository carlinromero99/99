package main

import "log"

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

func (p *Person) GetName() string {
	return p.name
}

func main() {
	person := NewPerson("Carlos")
	log.Printf("hello %s!", person.GetName())
}
