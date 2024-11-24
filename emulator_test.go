package main

import (
	"io"
	"os"
	"testing"
)
import "github.com/stretchr/testify/assert"

func Test_EmulatorLoadFont(t *testing.T) {
	// Given
	emulator, err := NewEmulator(NewTerminalDisplay())
	assert.NoError(t, err)

	// Then
	assert.Equal(t, emulator.memory[0xF0:0xF0+len(FontSprites)], FontSprites)
}

func Test_EmulatorLoadRom(t *testing.T) {
	// Setup
	fh, err := os.Open("roms/ibm-logo.ch8")
	assert.NoError(t, err)
	data, err := io.ReadAll(fh)
	assert.NoError(t, err)

	emulator, err := NewEmulator(NewTerminalDisplay())
	assert.NoError(t, err)

	// Test
	err = emulator.loadROM("roms/ibm-logo.ch8")
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, emulator.memory[0x200:0x200+len(data)], data)
}
