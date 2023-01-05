package main

import "fmt"

// Open-closed principle
// That is, open for extension, closed for modification

type Color int

const (
    red Color = iota
    green
    blue
)

type Size int

const (
    small Size = iota
    medium
    large
)

type Product struct {
    name string
    color Color
    size Size
}

type Filter struct {
    // filter settings
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.color == color {
            result = append(result, &products[i]) 
        }   
    }
    return result
}

/*
The later addition of the following methods violates the OCP, because it 
adds new behaviours to a structure which was previously created and tested.

The ideal implementation should be open to extensions, without modifying
things that have already been written and tested.
*/
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.size == size {
            result = append(result, &products[i]) 
        }   
    }
    return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.size == size && v.color == color {
            result = append(result, &products[i]) 
        }   
    }
    return result
}

/*
Solution: Specification pattern
*/

// Specification interface's goal is to test wheter or not a product satisfies some criteria.
// In this case, the interface type is open to extension, but closed to modification.
type Specification interface {
    IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
    color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
   return p.color == c.color 
}

type SizeSpecification struct {
    size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
   return p.size == s.size
}

 type AndSpecification struct {
    first, second Specification
 }

 func (a AndSpecification) IsSatisfied(p *Product) bool {
    return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
 }

type BetterFilter struct {}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if spec.IsSatisfied(&v) {
            result = append(result, &products[i]) 
        }   
    }
    return result
}


func main() {
    apple := Product{"Apple", green, small}
    tree := Product{"Tree", green, large}
    house := Product{"House", blue, large}
    products := []Product{apple, tree, house}
    f := Filter{}

    fmt.Println("Green products (old):")
    for _, v := range f.FilterByColor(products, green) {
        fmt.Printf(" - %s is green\n", v.name)
    }

    // Using Specification pattern:
    greenSpec := ColorSpecification{green}
    bf := BetterFilter{}

    fmt.Println("Green products (new):")
    for _, v := range bf.Filter(products, greenSpec) {
        fmt.Printf(" - %s is green\n", v.name)
    }
    
    largeSpec := SizeSpecification{large}
    lgSpec := AndSpecification{greenSpec, largeSpec}

    fmt.Println("Green and large products:")
    for _, v := range bf.Filter(products, lgSpec) {
        fmt.Printf(" - %s is green and large\n", v.name)
    }
}
