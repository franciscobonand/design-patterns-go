package main

import "fmt"

// Dependecy Inversion Principle
// High-level modules should not depend on low-level modules, both should
// depend on abstractions

type Relationship int

const (
    Parent Relationship = iota
    Child
    Sibling
)

type Person struct {
    name string
    //
}

type Info struct {
    from *Person
    relationship Relationship
    to *Person
}

// Low-level module (storage)
type Relationships struct {
    relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
    r.relations = append(r.relations, Info{parent, Parent, child})
    r.relations = append(r.relations, Info{child, Child, parent})
}

/*
High-level module (operate on data)

This is breaking the dependency inversion problem as it's using the low-level
module directly. Because of that, many problems can arise, such as the low-level
module changing the way it stores data (replacing a slice with a database, for
example).   
*/
type Research struct {
    // breaking DIP:
    relationships Relationships
}

func (r *Research) Investigate() {
    relations := r.relationships.relations
    for _, rel := range relations {
        if rel.from.name == "John" && rel.relationship == Parent {
            fmt.Println("John has a child called", rel.to.name)
        }
    }
}

// Solution: using abstractions

// Low-level module implements its own means to search
type RelationshipBrowser interface {
    FindAllChildrenOf(name string) []*Person
}

// As we're on the low-level module, it's ok to depend on direct access to the storage
func (r *Relationships) FindAllChildrenOf(name string) []*Person {
    result := make([]*Person, 0)
    for _, rel := range r.relations {
        if rel.from.name == name && rel.relationship == Parent {
            result = append(result, rel.to)
        }
    }
    return result
}

// High-level module is implemented using an abstraction
type Research2 struct {
    browser RelationshipBrowser
}

func (r *Research2) Investigate() {
    for _, p := range r.browser.FindAllChildrenOf("John") {
        fmt.Println("John has a child called", p.name)
    }
}

func mainDIP() {
    parent := Person{"John"}
    child1 := Person{"Maria"}
    child2 := Person{"Mark"}

    relationships := Relationships{}
    relationships.AddParentAndChild(&parent, &child1)
    relationships.AddParentAndChild(&parent, &child2)

    r := Research{relationships}
    r.Investigate()


    // Using abstractions
    r2 := Research2{&relationships}
    r2.Investigate()
}
