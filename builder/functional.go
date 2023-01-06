package main

import "fmt"

type PersonF struct {
    name, position string
}

// The benefit of this implementation is that it's easy to extend the builder
// with new build options without creating new builders which will eventually be aggregated
type personMod func(*PersonF)
type PersonBuilderF struct {
   actions []personMod 
}

func (b *PersonBuilderF) Called(name string) *PersonBuilderF {
    b.actions = append(b.actions, func(p *PersonF) {
        p.name = name
    })
    return b
}

func (b *PersonBuilderF) Build() *PersonF {
    p := PersonF{}
    for _, a := range b.actions {
        a(&p)

    }
    return &p
}

// Extending the builder:
func (b *PersonBuilderF) WorksAdA(position string) *PersonBuilderF {
    b.actions = append(b.actions, func(p *PersonF) {
        p.position = position
    })
    return b
}

func main() {
    b := PersonBuilderF{}
    p := b.Called("Chico").WorksAdA("Mimic").Build()
    fmt.Println(*p)
}
