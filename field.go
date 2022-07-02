package main

import (
	"fmt"
)

const (
	pwdHeight = 1.0
	pwdWidth = 4.0

	txtHeight = 1.0
	txtWidth = 4.0
)

type Field interface {
	// Return string representation of value
	String() string

	// Parse from returned string (from user)
	Parse(spec string)

	// GetLabel must return Label of Field (if not input return "")
	GetLabel() string

	// GetName returns name of field (if applicable)
	GetName() string

	// generate Formspec entry at Offset
	Formspec(Offset) (spec string, size Offset)
}

/*`pwdfield[<X>,<Y>;<W>,<H>;<name>;<label>]`*/
type Passwd struct {
	Name, Label string

	Content string
}

func (p Passwd) String() string {
	return p.Content
}

func (p Passwd) GetLabel() string {
	return p.Label
}

func (p Passwd) GetName() string {
	return p.Name
}

func (p Passwd) Parse(spec string) {
	p.Content = spec
}

func (p Passwd) Formspec(off Offset) (string, Offset) {
	return fmt.Sprintf("pwdfield[%f,%f;%f,%f;%s;%s]\n", off[X], off[Y], pwdWidth, pwdHeight, p.Name, p.Label), Offset{pwdWidth, pwdHeight}
}

/*`field[<X>,<Y>;<W>,<H>;<name>;<label>;<default>]`*/
type TextField struct {
	Name, Label, Default string

	Content string
}

func (p TextField) String() string {
	return p.Content
}

func (p TextField) GetLabel() string {
	return p.Label
}

func (p TextField) GetName() string {
	return p.Name
}

func (p TextField) Parse(spec string) {
	p.Content = spec
}

func (p TextField) Formspec(off Offset) (string, Offset) {
	return fmt.Sprintf("field[%f,%f;%f,%f;%s;%s;%s]\n", off[X], off[Y], txtWidth, txtHeight, p.Name, p.Label, p.Default), Offset{txtWidth, txtHeight}
}

/*`textarea[<X>,<Y>;<W>,<H>;<name>;<label>;<default>]`*/
type TextArea struct {
	Name, Label, Default string

	Size Offset

	Content string
}

func (p TextArea) String() string {
	return p.Content
}

func (p TextArea) GetLabel() string {
	return p.Label
}

func (p TextArea) GetName() string {
	return p.Name
}

func (p TextArea) Parse(spec string) {
	p.Content = spec
}

func (p TextArea) Formspec(off Offset) (string, Offset) {
	return fmt.Sprintf("textarea[%f,%f;%f,%f;%s;%s;%s]\n", off[X], off[Y], p.Size[X], p.Size[Y], p.Name, p.Label, p.Default), p.Size
}
