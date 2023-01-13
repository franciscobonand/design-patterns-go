package main

import "fmt"

type Employee struct {
    Name, Position string
    AnnualIncome int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
    return func(name string) *Employee {
        return &Employee{name, position, annualIncome}
    }
}

// non-functional approach
type EmployeeFactory struct {
    Position string
    AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
    return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
    return &EmployeeFactory{position, annualIncome}
}

func mainG() {
    // functional approach usage
    // The following functions can be passed as arguments to other functions,
    // or also stored in map structures
    developerFactory := NewEmployeeFactory("developer", 60000)
    managerFactory := NewEmployeeFactory("manager", 80000)
    
    developer := developerFactory("Adam")
    manager := managerFactory("Jane")

    fmt.Println(developer)
    fmt.Println(manager)

    // non-functional approach usage
    // The only advantage of this approach is the possibility of changing the
    // factory values after it was created. This option is not the most idiomatic one
    bossFactory := NewEmployeeFactory2("CEO", 12332112)
    boss := bossFactory.Create("Robert")
    fmt.Println(boss)
}
