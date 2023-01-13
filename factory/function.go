package main

import "fmt"

type person struct {
    name string
    age int
    eyeCount int // it'd be useful to have a default value for this field
}

type tiredPerson struct {
    name string
    age int
}

/*
An interface can be used alongside the factory function. This is useful to
encapsulate the structure information.
In this case, for example, two different underlying structures are being used,
depending on the user input to the constructor method 'NewPerson'.
*/
type Person interface {
    SayHello()
}

func (p *person) SayHello() {
    fmt.Printf("Hey, my name is %s, I am %d years old\n", p.name, p.age)
}

func (tp *tiredPerson) SayHello() {
    fmt.Println("Sorry I'm too tired to talk")
}

// NewPerson is a factory function that returns a Person interface
// with a default value for the 'eyeCount' property
func NewPerson(name string, age int) Person {
    if age > 80 {
        return &tiredPerson{name, age}
    }
    return &person{name, age, 2}
}

func mainFF() {
    p := NewPerson("John", 33) 
    p.SayHello()
    tp := NewPerson("Maria", 100)
    tp.SayHello()
}
