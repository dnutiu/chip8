package main

import (
	"fmt"
)

// DisplayWidth represents the display's width in pixels.
const DisplayWidth = 64

// DisplayHeight represents the display's height in pixels.
const DisplayHeight = 32

// Display interface defines methods for drawing and redrawing the display.
type Display interface {
	Redraw()
	Draw()
}

// TerminalDisplay models the Chip8's display.
type TerminalDisplay struct {
	// displayData holds the display data, each bool corresponds to a pixel.
	displayData [DisplayWidth * DisplayHeight]bool
}

// NewTerminalDisplay creates and returns a new TerminalDisplay.
func NewTerminalDisplay() *TerminalDisplay {
	return &TerminalDisplay{
		displayData: [DisplayWidth * DisplayHeight]bool{},
	}
}

// Redraw implements the Display interface for TerminalDisplay.
// It clears the terminal screen.
func (td *TerminalDisplay) Redraw() {
	// ANSI Escape code to clear the screen and move cursor to top left
	fmt.Printf("\033[2J\033[1;1H")
}

// Draw implements the Display interface for TerminalDisplay.
// It prints the current state of the display to the terminal.
func (td *TerminalDisplay) Draw() {
	for row := 0; row < DisplayHeight; row++ {
		for column := 0; column < DisplayWidth; column++ {
			if td.displayData[row*DisplayWidth+column] {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
