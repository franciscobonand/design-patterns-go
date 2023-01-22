package main

import "fmt"

// Shapes: circle, square
// Displays: raster, vector
// Result: RasterCircle, VectorCircle, ...

// Solution: Bridge pattern!

type Renderer interface {
    RenderCircle(radius float32)
    // RenderTriangle...
    // RenderSquare...
}

// Vector mode definition
type VectorRenderer struct {
    // Utility information on how vectors are constructed
}

func (v *VectorRenderer) RenderCircle(radius float32) {
    fmt.Println("Drawing a circle of radius", radius)
}

// Raster mode definition
type RasterRenderer struct {
    // Utility information on how raster images are constructed
}

func (r *RasterRenderer) RenderCircle(radius float32) {
    fmt.Println("Drawing pixels for circle of radius", radius)
}

// Circle shape definition
type Circle struct {
    renderer Renderer
    radius float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
    return &Circle{
        renderer,
        radius,
    }
}

func (c *Circle) Draw() {
    c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
    c.radius *= factor
}

func main() {
    // raster := RasterRenderer{}
    vector := VectorRenderer{}
    circle := NewCircle(&vector, 5)
    circle.Draw()
    circle.Resize(2)
    circle.Draw()
}
