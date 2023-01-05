package main

import (
	"fmt"
)

// Liskov Substitution Principle
// Deals with inheritance, not really applicable in Go


type Sized interface {
    GetWidth() int
    SetWidth(width int)
    GetHeight() int
    SetHeight(height int)
}


type Rectangle struct {
    width, height int
}

func (r *Rectangle) GetWidth() int {
    return r.width
}

func (r *Rectangle) SetWidth(width int) {
    r.width = width
}

func (r *Rectangle) GetHeight() int {
    return r.height
}

func (r *Rectangle) SetHeight(height int) {
    r.height = height
}

/*
The following structure Square breaks the Liskov Substitution Principle.
That's because structures extended from another one which is higher in the hierarchy
are still expected to behave the same way.

In this case, Square is extended from Rectangle, but we can see that the function
'UseIt' behaves differently depending on the structure (for Squares it doesn't
work as expected).
*/
type Square struct {
    Rectangle
}

func NewSquare(size int) *Square{
    sq := Square{}
    sq.width = size
    sq.height = size
    return &sq
}

func (s *Square) SetWidth(width int) {
    s.width = width
    s.height = width
}

func (s *Square) SetHeight(height int) {
    s.width = height
    s.height = height
}

func UseIt(sized Sized) {
    width := sized.GetWidth()
    sized.SetHeight(10)
    expectedArea := 10 * width
    actualArea := sized.GetHeight() * sized.GetWidth()

    fmt.Println("Expected an area of", expectedArea, ", but got", actualArea)
}

// Possible solution for Square violation:
type Square2 struct {
    size int // width, height
}

func (s *Square2) Rectangle() Rectangle {
    return Rectangle{s.size, s.size}
}

func mainLSP() {
    r := &Rectangle{2,3}
    UseIt(r)
    s := NewSquare(2)
    UseIt(s)
}
