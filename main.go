package main

import "log"

type Person struct {
	name string
	age  int64
}

func NewPerson(name string, age int64) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() int64 {
	return p.age
}

func main() {
	person := NewPerson("Carlos", 29)
	log.Printf("hello %s!, you are %d", person.GetName(), person.GetAge())
}
