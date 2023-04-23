package main

import "fmt"

type iNode interface {
	print(string)
	clone() iNode
}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() iNode {
	return &file{name: f.name + "_clone"}
}

type folder struct {
	children []iNode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() iNode {
	cloneFolder := &folder{name: f.name + "_clone"}

	var tempChildren []iNode
	tempChildren = append(tempChildren, f.children...)
	cloneFolder.children = tempChildren

	return cloneFolder
}
