package main

import "fmt"

type Buffer struct {
    width, height int
    buffer []rune
}

func NewBuffer(width, height int) *Buffer {
    return &Buffer {
        width: width,
        height: height,
        buffer: make([]rune, width * height),
    }
}

func (b *Buffer) At(index int) rune {
    return b.buffer[index]
}

type Viewport struct {
    buffer *Buffer
    offset int
}

func NewViewport(buffer *Buffer) *Viewport {
    return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharAt(index int) rune {
    return v.buffer.At(v.offset + index)
}


// Console is a Façade that makes it easy to handle multiple buffers and viewports
type Console struct {
    buffers []*Buffer
    viewports []*Viewport
    offset int
}

func NewDefaultConsole() *Console {
    b := NewBuffer(200, 150)
    v := NewViewport(b)
    return &Console{
        buffers: []*Buffer{b},
        viewports: []*Viewport{v},
        offset: 0,
    }
}

func (c *Console) GetCharAt(index int) rune {
    return c.viewports[0].GetCharAt(index)
}

func main() {
    // Using the Console Façade makes it much simpler to handle the more complicated
    // structure which is composed by multiple viewports and buffers
    c := NewDefaultConsole()
    u := c.GetCharAt(1)
    fmt.Println(u)
}
