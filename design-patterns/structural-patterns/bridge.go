package main

import "fmt"

//say you have two classes shapes and colors
// you add two shapes -> square and circle
// and two colors -> red and blue
// total combinations are 4, and this can grow exponentially
// bridge pattern aims to solve for this

//go example : two computers Mac and Windows
// two printers Epson and HP -> need to work together in any combination
// client doesn't want to worry about details of connecting printers to computers

// we create two hierarchies
// abstraction hierarchy and implementation hierarchy
// both will communicate each other via a bridge
// both the abstraction and implementation can be developed independently

type Computer2 interface {
	Print()
	SetPrinter(Printer)
}
type Mac2 struct {
	printer Printer
}

func (m *Mac2) Print() {
	fmt.Println("print request for mac")
	m.printer.PrintFile()
}

func (m *Mac2) SetPrinter(p Printer) {
	m.printer = p
}

type Windows2 struct {
	printer Printer
}

func (w *Windows2) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows2) SetPrinter(p Printer) {
	w.printer = p
}

type Printer interface {
	PrintFile()
}

// implementation for printers
type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("printing by an EPSON printer")
}

type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac2{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	//similarly can be done for windows computer as well

	// very useful i'd say
}
