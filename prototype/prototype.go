package main

import "fmt"

type Address struct {
    Street, City, Country string
}

type Person struct {
    Name string
    Address *Address
}

func mainP() {
    // Creating a Person
    john := Person{
        Name: "John",
        Address: &Address{
            Street: "London Road",
            City: "London",
            Country: "England",
        },
    }
    // Trying to copy a person
    // jane := john
    // jane.Name = "Jane" // ok
    // jane.Address.Street = "Baker Street" // not ok

    // As the pointer is being copied, modifications on Jane's address will also
    // modify John's address (because the pointer references the same underlying structure)
    // fmt.Println(john, john.Address)
    // fmt.Println(jane, jane.Address)

    // Deep copying a person
    janedc := john
    janedc.Address = &Address{
        john.Address.Street,
        john.Address.City,
        john.Address.Country,
    }
    janedc.Address.Street = "Baker Street"
    fmt.Println(john, john.Address)
    fmt.Println(janedc, janedc.Address)

    // As it can be seen, deep copying the structure this way is not scalable
}
