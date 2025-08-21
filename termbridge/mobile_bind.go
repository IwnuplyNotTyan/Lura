package termbridge

// This file defines gomobile-friendly types and shims with only bindable signatures.

// Output is the bindable interface for receiving lines from the game.
type Output interface {
    OnLine(line string)
}

// Mobile is the bindable wrapper exposed to Android/iOS via gomobile bind.
type Mobile struct {
    cli *MobileCLI
}

// NewMobile creates a Mobile that forwards output to the provided Output.
func NewMobile(output Output) *Mobile {
    m := &Mobile{cli: NewMobileCLI(output)}
    return m
}

// Start starts the bridged CLI.
func (m *Mobile) Start() error { return m.cli.Start() }

// SendLine sends a line of input.
func (m *Mobile) SendLine(line string) error { return m.cli.SendLine(line) }

// Close stops the CLI and restores stdio.
func (m *Mobile) Close() { m.cli.Close() }

