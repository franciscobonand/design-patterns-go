package main

import "fmt"

// Person has two distinct sets of information, and you might want to have a
// separate builder for each one.
type Person struct {
    // address
    StreetAddress, Postcode, City string

    // job
    CompanyName, Position string
    AnnualIncome int
}

// PersonBuilder builds a person entity
type PersonBuilder struct {
    person *Person
}

func NewPersonBuilder() *PersonBuilder {
    return &PersonBuilder{ &Person{} }
}

func (b *PersonBuilder) Lives() *AddressBuilder {
    return  &AddressBuilder{*b}
}

func (b *PersonBuilder) Works() *JobBuilder {
    return  &JobBuilder{*b}
}

// AddressBuilder is a PersonBuilder used for address info
type AddressBuilder struct {
    PersonBuilder
}

func (ab *AddressBuilder) At(streetAddr string) *AddressBuilder {
    ab.person.StreetAddress = streetAddr
    return ab
}

func (ab *AddressBuilder) In(city string) *AddressBuilder {
    ab.person.City = city
    return ab
}

func (ab *AddressBuilder) WithPostcode(postcode string) *AddressBuilder {
    ab.person.Postcode = postcode
    return ab
}

// JobBuilder is a PersonBuilder used for job info
type JobBuilder struct {
    PersonBuilder
}

func (jb *JobBuilder) At(companyName string) *JobBuilder {
    jb.person.CompanyName = companyName
    return jb
}

func (jb *JobBuilder) AsA(position string) *JobBuilder {
    jb.person.Position = position
    return jb
}

func (jb *JobBuilder) Earning(income int) *JobBuilder {
    jb.person.AnnualIncome = income
    return jb
}

// Build "bundles" all of a Persons definitions and returns it
func (b *PersonBuilder) Build() *Person {
    return b.person
}


func mainBF() {
    pb := NewPersonBuilder()
    pb.
        Lives().
            At("Av. Amazonas 1009").
            In("Rio de Janeiro").
            WithPostcode("64423102").
        Works().
            At("Globo").
            AsA("Producer").
            Earning(1234)
    person := pb.Build()
    fmt.Println(person)
}
