package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Here are some better approaches for deep copying a structure

type Person2 struct {
    Name string
    Address *Address
    Friends []string
}

// Creating DeepCopy methods for each structure
func (a *Address) DeepCopy() *Address {
    return &Address{
        a.Street,
        a.City,
        a.Country,
    }
}

func (p *Person2) DeepCopy() *Person2 {
    q := *p
    q.Address = p.Address.DeepCopy()
    copy(q.Friends, p.Friends)
    return &q
}


// Using serialization (better approach)
func (p *Person2) DeepCopyS() *Person2 {
    b := bytes.Buffer{}
    e := gob.NewEncoder(&b)
    _ = e.Encode(p)

    d := gob.NewDecoder(&b)
    result := Person2{}
    _ = d.Decode(&result)
    return &result
}

func mainC() {
    john := Person2{
        Name: "John",
        Address: &Address{
            Street: "London Road",
            City: "London",
            Country: "England",
        },
        Friends: []string{"Matt", "Maria", "Ronald"},
    }

    jane := john.DeepCopyS()
    jane.Name = "Jane"
    jane.Address.Street = "Baker Street"
    jane.Friends = append(jane.Friends, "Angela")

    fmt.Println(john, john.Address)
    fmt.Println(jane, jane.Address)
}
