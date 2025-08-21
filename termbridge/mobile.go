package termbridge

import (
    "errors"
)

// MobileCLI is a gomobile-bindable facade to run an unchanged CLI game under bridged stdio.
// Workflow:
// 1) Your game code calls termbridge.RegisterEntrypoint(func() { /* start your game */ }) once (e.g., in init()).
// 2) Android side creates MobileCLI with an OutputHandler, calls Start(), then SendLine() for user input, and Close() to stop.
type MobileCLI struct {
    bridge *Bridge
}

// NewMobileCLI creates a new instance. Provide an OutputHandler to receive game output lines.
func NewMobileCLI(output OutputHandler) *MobileCLI {
    return &MobileCLI{bridge: NewBridge(output)}
}

// registeredEntrypoint is a package-level function the game should register.
var registeredEntrypoint func()

// RegisterEntrypoint sets the function to run when MobileCLI.Start is called.
// Your CLI game's init() can call this with the function that starts the game loop.
// The function should block until the game exits (like a normal main would).
func RegisterEntrypoint(fn func()) {
    registeredEntrypoint = fn
}

// Start activates stdio bridging and runs the registered entrypoint in a new goroutine.
func (m *MobileCLI) Start() error {
    if m.bridge == nil {
        return errors.New("nil bridge")
    }
    if err := m.bridge.Start(); err != nil {
        return err
    }
    go func() {
        if registeredEntrypoint != nil {
            registeredEntrypoint()
        } else {
            // Inform the client that no entrypoint was registered
            _ = m.bridge.SendLine("exit")
        }
    }()
    return nil
}

// SendLine sends a single line of input to the game's stdin.
func (m *MobileCLI) SendLine(line string) error {
    return m.bridge.SendLine(line)
}

// Close stops the game I/O bridge and restores stdio.
func (m *MobileCLI) Close() {
    if m.bridge != nil {
        m.bridge.Close()
    }
}

