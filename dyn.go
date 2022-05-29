package main

import (
	"fmt"
	"strings"
)

type DynContent interface {
	Evaluate(string, *SignPos) string
}

type PlayerCnt struct {
	Srv string
	Pad int
}

func (pc *PlayerCnt) Evaluate(text string, pos *SignPos) string {
	return fmt.Sprintf("%d", playersOnSrv(pc.Srv))
}

type Center struct {
	Line   int
	Length int
	Rune   rune
	Sub    int
}

func (ce *Center) Evaluate(text string, pos *SignPos) string {
	p := (ce.Length - len(strings.Split(text, "\n")[ce.Line-1])) / 2

	return strings.Repeat(string(ce.Rune), p)
}
