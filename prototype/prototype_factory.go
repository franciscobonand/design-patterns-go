package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// This approach is more convenient to use 

type Address2 struct {
    Suite int
    Street, City string
}

type Employee struct {
    Name string
    Office Address2
}

func (p *Employee) DeepCopy() *Employee {
    b := bytes.Buffer{}
    e := gob.NewEncoder(&b)
    _ = e.Encode(p)

    d := gob.NewDecoder(&b)
    result := Employee{}
    _ = d.Decode(&result)
    return &result
}

// Pre-defined offices (prototypes)
var mainOffice = Employee{
    Office: Address2{
        Suite: 0,
        Street: "Av. Amazonas",
        City: "Manaus",
    },
}

var auxOffice = Employee{
    Office: Address2{
        Suite: 0,
        Street: "Rua São João",
        City: "Fortaleza",
    },
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
    result := proto.DeepCopy()
    result.Name = name
    result.Office.Suite = suite
    return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
    return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
    return newEmployee(&auxOffice, name, suite)
}

func main() {
    john := NewMainOfficeEmployee("John", 100)
    jane := NewAuxOfficeEmployee("Jane", 200)

    fmt.Println(john)
    fmt.Println(jane)
}
