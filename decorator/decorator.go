package main

import "fmt"

type Shape interface {
    Render() string
}

type Circle struct {
    Radius float32
}

func (c *Circle) Render() string {
    return fmt.Sprintf("Circle of radius %.2f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
    c.Radius *= factor
}

type Square struct {
    Side float32
}

func (s *Square) Render() string {
    return fmt.Sprintf("Square with side %.2f", s.Side)
}


type ColoredShape struct {
    Shape Shape
    Color string
}

func (c *ColoredShape) Render() string {
    return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}


type TransparentShape struct {
    Shape Shape
    Transparency float32
}

func (t *TransparentShape) Render() string {
    return fmt.Sprintf("%s has %d%% transparecy", t.Shape.Render(), int(t.Transparency * 100.0))
}

func main() {
    circle := Circle{2}
    circle.Resize(2)
    fmt.Println(circle.Render())
    
    redCircle := ColoredShape{&circle, "Red"}
    fmt.Println(redCircle.Render())

    // redCircle.Resize(...) is unavailable as it's exclusive to the Circle struct,
    // this is a downside of decorators.
    // Upside: decorators can be composed!

    rhtCircle := TransparentShape{&redCircle, 0.5}
    fmt.Println(rhtCircle.Render())
}

