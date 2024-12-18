package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	MemorySize    = 4096
	RegistersSize = 16
)

// FontSprites are the sprites that hold the emulator.
var FontSprites = []byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

// Emulator emulates the Chip8 CPU.
type Emulator struct {
	// memory represents the emulator's RAM memory.
	memory [MemorySize]uint8
	// registers holds the general purpose registers.
	registers [RegistersSize]uint8
	// The programCounter register tracks the currently executing instruction
	programCounter uint16
	// The delayTimer register. It is decremented at a rate of 60 Hz until it reaches 0.
	delayTimer uint8
	// The soundTimer register. It is decremented at a rate of 60 Hz until it reaches 0.
	// It plays a beeping sound when it's value is different from 0.
	soundTimer uint8
	// The stackPointer register.
	stackPointer uint8
	// display held a Display instance.
	display Display
}

// NewEmulator creates a new emulator instance.
func NewEmulator(display Display) (*Emulator, error) {
	var emulator = &Emulator{
		memory:    [MemorySize]uint8{},
		registers: [RegistersSize]uint8{},
		display:   display,
	}

	emulator.loadFontData()

	return emulator, nil
}

// Emulate starts the emulation of the ROM specified at `path`.
func (e *Emulator) Emulate(path string) error {
	if err := e.loadROM(path); err != nil {
		return err
	}
	return nil
}

// loadROM loads the ROM found at the romFile path into the emulator's RAM memory.
func (e *Emulator) loadROM(romFile string) error {
	file, err := os.Open(romFile)
	if err != nil {
		return fmt.Errorf("error opening ROM file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Failed to close file", slog.String("err", err.Error()))
		}
	}(file)

	// Check ROM length if it overflows max RAM size.
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting file info: %v", err)
	}
	romSize := stat.Size()
	slog.Info(fmt.Sprintf("Open ROM %s of size %d bytes.", romFile, romSize))
	if romSize > MemorySize-0x200 {
		return fmt.Errorf("ROM at %s overflows emulator's RAM size of 4kB", romFile)
	}

	// Read the ROM into memory
	romData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading ROM data: %v", err)
	}
	copy(e.memory[0x200:], romData)

	e.debugPrintMemory()
	return nil
}

// loadFontData loads the font data into RAM.
func (e *Emulator) loadFontData() {
	slog.Info("Loading font data...")
	for i, sprite := range FontSprites {
		e.memory[0xF0+i] = sprite
	}
}

// debugPrintMemory prints memory inside the console when debug is activated
func (e *Emulator) debugPrintMemory() {
	if programLevel.Level() == slog.LevelDebug {
		for start := 0; start < len(e.memory); start += 16 {
			end := start + 16
			slog.Debug(fmt.Sprintf("Memory[%x:%x]= %#v", start, end, e.memory[start:end]))
		}
	}
}
