package core

import (
	"os"

	"golang.org/x/term"
)

const (
	MAX_DISPLAY_WIDTH = 72
	MIN_DISPLAY_WIDTH = 42
)

type TermContext struct {
	Width      int
	Height     int
	UsedWidth  int
	IsTerminal bool
}

func NewTermContext() *TermContext {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return &TermContext{0, 0, MAX_DISPLAY_WIDTH, false}
	}

	usedWidth := max(MIN_DISPLAY_WIDTH, min(w, MAX_DISPLAY_WIDTH))

	return &TermContext{w, h, usedWidth, true}
}
