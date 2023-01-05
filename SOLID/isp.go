package main

// Interface Segregation Principle
// Basically, don't put too much in an interface ('God class')

type Document struct {

}

type Machine interface {
    Print(d Document)
    Fax(d Document)
    Scan(d Document)
}

// MultiFunctionPrinter is able to print, fax and scan, so all methods make sense
type MultiFunctionPrinter struct {}

func (m *MultiFunctionPrinter) Print(d Document) {
    // implementation
}

func (m *MultiFunctionPrinter) Fax(d Document) {
    // implementation
}

func (m *MultiFunctionPrinter) Scan(d Document) {
    // implementation
}


// OldPrinter can only print, but all the methods are implemented so it is compliant
// with some API. This is a problem created by putting too much into just one interface.
type OldPrinter struct {}

func (p *OldPrinter) Print(d Document) {
    // implementation
}

func (p *OldPrinter) Fax(d Document) {
    // write panic message?
}

func (p *OldPrinter) Scan(d Document) {
    // write panic message?
}


// Solution: Interface Segregation
type Printer interface {
    Print(d Document)
}

type Scanner interface {
    Scan(d Document)
}

// MyPrinter can only print
type MyPrinter struct {}

func (m *MyPrinter) Print(d Document) {
    // implementation
}

// Photocopier can print and scan
type Photocopier struct {}

func (p *Photocopier) Print(d Document) {
    // implementation
}

func (p *Photocopier) Scan(d Document) {
    // implementation
}

// Interfaces can be composed:
type MultiFunctionDevice interface {
    Printer
    Scanner
    // Fax
}

// Using decorators to create multi-function structures:
type MultiFunctionMachine struct {
    printer Printer
    scanner Scanner
}

func (m *MultiFunctionMachine) Print(d Document) {
    m.printer.Print(d)
}

func (m *MultiFunctionMachine) Scan(d Document) {
    m.scanner.Scan(d)
}

func mainISP() {

}
