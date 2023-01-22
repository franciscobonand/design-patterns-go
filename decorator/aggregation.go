package main

import "fmt"

/* First attempt defining a Dragon, which is both a Bird and a Lizard

type Bird struct {
    Age int
}

func (b *Bird) Fly() {
    if b.Age >= 10 {
        fmt.Println("Flying!")
    }
}

type Lizard struct {
    Age int
}

func (l *Lizard) Crawl() {
    if l.Age < 10 {
        fmt.Println("Crawling!")
    }
}

// A Dragon is both a Lizard and a Bird. It could be defined as the
// aggregation of the two structs.
type Dragon struct {
    Bird
    Lizard
}

// Age and SetAge methods try to mitigate the ambiguous selector problem
func (d *Dragon) Age() int {
    return d.Bird.Age
}

func (d *Dragon) SetAge(age int) {
    d.Bird.Age = age
    d.Lizard.Age = age
}
*/

// Second attempt, using the decorator pattern:

// Aged defines a contract of getter and setter for 'age' field in structs
type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
    if b.age >= 10 {
        fmt.Println("Flying!")
    }
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
    if l.age < 10 {
        fmt.Println("Crawling!")
    }
}

// The Dragon structure will define methods which are proxy to
// the Bird o Lizard methods
type Dragon struct {
    bird Bird
    lizard Lizard
}

func (d *Dragon) Age() int {
    return d.bird.Age()
}

func (d *Dragon) SetAge(age int) {
    d.bird.SetAge(age)
    d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
    d.bird.Fly()
}

func (d *Dragon) Crawl() {
    d.lizard.Crawl()
}

func NewDragon() *Dragon {
    return &Dragon{Bird{}, Lizard{}}
}

func mainA() {
    /* Using first attempt's structures:
        d := Dragon{}
        // Problem: setting dragon Age (ambiguous selector)
        // d.Age = 10
        d.SetAge(5)
        // Despite the SetAge method, it's still possible to alter one of the ages
        d.Bird.Age = 55
        d.Fly()
        d.Crawl()
    */

    // Using the decorator pattern:
    d := NewDragon()
    d.SetAge(5)
    d.Fly()
    d.Crawl()
}
