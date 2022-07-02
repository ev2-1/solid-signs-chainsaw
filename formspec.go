package main

import (
	"strings"
)

// Escape... escapes strings for formspecs
func Escape(str string) string {
	str = strings.ReplaceAll(str, "[", "\\[")
	str = strings.ReplaceAll(str, "]", "\\]")

	str = strings.ReplaceAll(str, "<", "\\<")
	str = strings.ReplaceAll(str, ">", "\\>")
	return str
}

const (
	X = 0
	Y = 1
)

type Offset [2]float32

func (a Offset) Add(b Offset) (o Offset) {
	for k := range o {
		o[k] = a[k] + b[k]
	}

	return
}
