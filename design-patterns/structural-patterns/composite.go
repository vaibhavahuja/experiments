package main

import "fmt"

//useful when there is a tree like structure
// for example you have a box and product
// each box can have multiple products and boxes
// if i ask you the total cost, nesting can be messy
// instead declare a method which calculates the cost of each box/product
// use that. Only used in case of tree like structures

//example File System
// Need to run a search for a particular keyword in your file system
// Search operation applies to both files and folders

type Component interface {
	search(string)
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Println("recursively searching for keyword %s in folder %s", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}

//more like a simple recursive search...
