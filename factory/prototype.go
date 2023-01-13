package main

import "fmt"


const (
    Developer = iota
    Manager
)

// In this approach, there are predefined options from which the user can choose
// to instantiate the structure
func NewEmployee3(role int) *Employee {
    switch role {
    case Developer:
        // instead of the structure itself, one could return an interface
        return &Employee{"", "developer", 60000}
    case Manager:
        return &Employee{"", "manager", 80000}
    default:
        panic("unsupported role")
    }
}

func main() {
    m := NewEmployee3(Manager)
    m.Name = "Bob"
    fmt.Println(m)
}
