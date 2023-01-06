package main

import (
	"fmt"
	"strings"
)

const (
    indentSize = 2
)

// Representation of an HTML method and how it's displayed (printed)
type HTMLElement struct {
    name, text string
    elements []HTMLElement
}

func (e *HTMLElement) String() string {
    return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
    sb := strings.Builder{}
    i := strings.Repeat(" ", indentSize * indent)
    sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

    if len(e.text) > 0 {
        sb.WriteString(strings.Repeat(" ", indentSize * (indent + 1)))
        sb.WriteString(e.text)
        sb.WriteString("\n")
    }

    for _, el := range e.elements {
        sb.WriteString(el.string(indent+1))
    }

    sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
    return sb.String()
}

// Builder definition for HTML
type HTMLBuilder struct {
    rootName string
    root HTMLElement
}

func NewHTMLBuilder(rootName string) *HTMLBuilder {
    return &HTMLBuilder{
        rootName: rootName,
        root: HTMLElement{
            name: rootName,
            text: "",
            elements: []HTMLElement{},
        },
    }
}

func (b *HTMLBuilder) String() string {
    return b.root.String()
}

func (b *HTMLBuilder) AddChild(name, text string) {
    e := HTMLElement{name, text, []HTMLElement{}}
    b.root.elements = append(b.root.elements, e)
}

// This method is chainable
func (b *HTMLBuilder) AddChildFluent(name, text string) *HTMLBuilder {
    e := HTMLElement{name, text, []HTMLElement{}}
    b.root.elements = append(b.root.elements, e)
    return b
}

func main() {
    // Using stdlib strings.Builder to create an HTML is pretty verbose:
    hello := "hello"
    sb := strings.Builder{}
    sb.WriteString("<p>")
    sb.WriteString(hello)
    sb.WriteString("</p>")
    fmt.Println(sb.String())

    words := []string{"hello", "world"}
    sb.Reset()
    // <ul><li>...</li><<li>...</li><li>...</li>ul>
    sb.WriteString("<ul>")
    for _, v := range words {
        sb.WriteString("<li>")
        sb.WriteString(v)
        sb.WriteString("</li>")
    }
    sb.WriteString("</ul>")
    fmt.Println(sb.String())

    // Using HTMLBuilder
    b := NewHTMLBuilder("ul")
    b.AddChild("li", "hello")
    b.AddChild("li", "world")
    fmt.Println(b.String())

    // Using HTMLBuilder Fluent method
    bf := NewHTMLBuilder("ul")
    bf.AddChildFluent("li", "hello").AddChildFluent("li", "world")
    fmt.Println(bf.String())
}
