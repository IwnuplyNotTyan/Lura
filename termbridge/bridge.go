package termbridge

import (
    "bufio"
    "errors"
    "io"
    "os"
    "strings"
    "sync"
)

// OutputHandler receives lines written by the game to stdout/stderr.
// It should be non-blocking; long operations must be performed asynchronously.
type OutputHandler interface {
    OnLine(line string)
}

// Bridge redirects process-wide stdio (os.Stdin, os.Stdout, os.Stderr)
// to in-memory pipes so a CLI game that uses standard input/output
// can run unchanged while we feed input and capture output.
//
// Warning: stdio redirection is process-wide; only run one Bridge at a time.
type Bridge struct {
    outputHandler OutputHandler

    originalStdin  *os.File
    originalStdout *os.File
    originalStderr *os.File

    stdinReader  *os.File
    stdinWriter  *os.File
    stdoutReader *os.File
    stdoutWriter *os.File
    stderrReader *os.File
    stderrWriter *os.File

    closeOnce sync.Once
    closed    chan struct{}

    wg sync.WaitGroup
}

// NewBridge constructs a Bridge. Call Start to activate stdio redirection.
func NewBridge(outputHandler OutputHandler) *Bridge {
    return &Bridge{
        outputHandler: outputHandler,
        closed:        make(chan struct{}),
    }
}

// Start replaces os.Stdin, os.Stdout, os.Stderr with pipe-backed files
// and starts background readers that forward output to the OutputHandler.
func (b *Bridge) Start() error {
    if b.outputHandler == nil {
        return errors.New("nil OutputHandler")
    }

    var err error
    b.originalStdin = os.Stdin
    b.originalStdout = os.Stdout
    b.originalStderr = os.Stderr

    // stdin: our writer feeds the game's reader
    b.stdinReader, b.stdinWriter, err = os.Pipe()
    if err != nil {
        return err
    }
    // stdout/stderr: game's writers feed our readers
    b.stdoutReader, b.stdoutWriter, err = os.Pipe()
    if err != nil {
        return err
    }
    b.stderrReader, b.stderrWriter, err = os.Pipe()
    if err != nil {
        return err
    }

    os.Stdin = b.stdinReader
    os.Stdout = b.stdoutWriter
    os.Stderr = b.stderrWriter

    // Forward stdout and stderr lines to handler
    b.wg.Add(2)
    go b.forwardPipe("stdout", b.stdoutReader)
    go b.forwardPipe("stderr", b.stderrReader)

    return nil
}

// forwardPipe reads lines (\n-terminated) and forwards them.
// It attempts to preserve partial lines by flushing on Close.
func (b *Bridge) forwardPipe(_ string, r *os.File) {
    defer b.wg.Done()
    reader := bufio.NewReader(r)
    var partial strings.Builder
    for {
        select {
        case <-b.closed:
            // Flush any remaining buffered data as one line
            if partial.Len() > 0 {
                b.outputHandler.OnLine(partial.String())
            }
            return
        default:
        }

        chunk, err := reader.ReadString('\n')
        if len(chunk) > 0 {
            // Trim only the trailing newline; keep other whitespace and ANSI
            if chunk[len(chunk)-1] == '\n' {
                line := partial.String() + strings.TrimRight(chunk, "\n")
                partial.Reset()
                b.outputHandler.OnLine(line)
            } else {
                partial.WriteString(chunk)
            }
        }
        if err != nil {
            if err == io.EOF {
                // Flush what we have
                if partial.Len() > 0 {
                    b.outputHandler.OnLine(partial.String())
                    partial.Reset()
                }
                return
            }
            // Non-EOF error: terminate this pipe
            return
        }
    }
}

// SendLine writes a line of input to the game's stdin, appending a newline if missing.
func (b *Bridge) SendLine(line string) error {
    if b.stdinWriter == nil {
        return errors.New("bridge not started")
    }
    if !strings.HasSuffix(line, "\n") {
        line += "\n"
    }
    _, err := io.WriteString(b.stdinWriter, line)
    return err
}

// SendBytes writes raw bytes to stdin as-is. Useful for control sequences.
func (b *Bridge) SendBytes(data []byte) (int, error) {
    if b.stdinWriter == nil {
        return 0, errors.New("bridge not started")
    }
    return b.stdinWriter.Write(data)
}

// Close restores original stdio and stops background goroutines.
func (b *Bridge) Close() {
    b.closeOnce.Do(func() {
        close(b.closed)

        // Restore stdio as early as possible
        if b.originalStdin != nil {
            os.Stdin = b.originalStdin
        }
        if b.originalStdout != nil {
            os.Stdout = b.originalStdout
        }
        if b.originalStderr != nil {
            os.Stderr = b.originalStderr
        }

        // Close pipe ends to unblock readers/writers
        if b.stdinWriter != nil {
            _ = b.stdinWriter.Close()
        }
        if b.stdinReader != nil {
            _ = b.stdinReader.Close()
        }
        if b.stdoutWriter != nil {
            _ = b.stdoutWriter.Close()
        }
        if b.stdoutReader != nil {
            _ = b.stdoutReader.Close()
        }
        if b.stderrWriter != nil {
            _ = b.stderrWriter.Close()
        }
        if b.stderrReader != nil {
            _ = b.stderrReader.Close()
        }

        b.wg.Wait()
    })
}

