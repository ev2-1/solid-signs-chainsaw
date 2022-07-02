package main

import (
	//	"github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"

	//	"strings"
	"fmt"
)

// Form is a plan for a formspec with input fields
type Form struct {
	Title  string  // displayed at top of form
	Name   string  // technical name used to communicate results
	Fields []Field // content
}

// converts Form into formspec string
func (f Form) Formspec() *mt.ToCltShowFormspec {
	offset := Offset{0.5, 0.5}
	var spec string

	var width float32 = 1

	// Title
	spec = fmt.Sprintf("label[%f,%f;%s]\n", offset[X], offset[Y], f.Title)
	offset = offset.Add(Offset{0, 1})

	// all the fields
	for _, f := range f.Fields {
		str, size := f.Formspec(offset)
		offset[Y] += size[Y]

		fmt.Println(width, size[X])
		if size[X] > width {
			width = size[X]
		}

		spec += str
	}

	// size
	spec = fmt.Sprintf("size[%f,%f]\n", width+0.5, offset[Y]) + spec

	fmt.Println(spec)

	return &mt.ToCltShowFormspec{
		Formname: f.Name,
		Formspec: spec,
	}
}
