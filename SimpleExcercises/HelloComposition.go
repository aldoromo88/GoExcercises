package main

import (
	"fmt"
)

func main() {

	var student = &Student{
		Person: &Person{"Aldo Romo", 29},
	}

	var goku = &Saiyan{
		Person: &Person{"Goku", 35},
		Power:  9001,
	}

	fmt.Println(student.Name)
	student.SayHi()
	goku.SayHi()
}

type Person struct {
	Name string
	Age  int
}

type Student struct {
	*Person
}

type Saiyan struct {
	*Person
	Power int
}

func (p *Person) SayHi() {
	fmt.Println("Hi my name is", p.Name, "and I'm", p.Age, "years old")
}

func (s *Saiyan) SayHi() {
	fmt.Println("Hi my name is", s.Name + ", I'm", s.Age, "years old and my power is over", s.Power)
}
