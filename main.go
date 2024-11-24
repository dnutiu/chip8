package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	// Initialize logging
	programLevel.Set(slog.LevelInfo)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})))

	// Initialize emulator
	emulator, err := NewEmulator(NewTerminalDisplay())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to instantiate emulator instance %v", err))
		os.Exit(1)
	}

	err = emulator.Emulate("./roms/ibm-logo.ch8")
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to emulate ROM: %v", err))
		os.Exit(1)
	}
}
