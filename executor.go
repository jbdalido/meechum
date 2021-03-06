package meechum

import (
	"fmt"
	"io"
	//"log"
	"os"
	"os/exec"
	"syscall"
	//	"path"
	//"strings"
)

type Executor struct {
	Binary  string
	Workdir string
	Buffer  map[string][]byte
	SOut    io.Writer
	SErr    io.Writer
}

func NewExecutor(binary, w string) (*Executor, error) {
	// Lookup the git binary
	b, err := exec.LookPath(binary)
	if err != nil {
		return nil, fmt.Errorf("%s binary not found", binary)
	}

	return &Executor{
		SOut:    os.Stdout,
		SErr:    os.Stderr,
		Binary:  b,
		Workdir: w,
	}, nil
}

func (e *Executor) SetOut(sout, serr io.Writer) {
	e.SOut = sout
	e.SErr = serr
}

// Git execute
func (e *Executor) Do(args []string) (string, error, int) {

	buffer := NewBufferizer()
	// Set buffers for this run
	stdout := io.MultiWriter(os.Stdout, buffer)
	stderr := io.MultiWriter(os.Stdout, buffer)

	// Setup work directory and command
	cmd := &exec.Cmd{
		Path:   e.Binary,
		Args:   append([]string{e.Binary}, args...),
		Stdout: stdout,
		Stderr: stderr,
	}

	// Get error from cmd, stdout/err from the buffer and the exit code
	// There is no way for me to say if this gonna work under a windows
	// Environment. Sorry ! PR are welcomes
	err := cmd.Run()
	std := buffer.Get()
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	if err != nil {
		return std, err, waitStatus.ExitStatus()
	}
	return std, nil, 0
}

type Bufferizer struct {
	Buff []byte
}

func NewBufferizer() *Bufferizer {
	return &Bufferizer{
		Buff: make([]byte, 100),
	}
}

func (b *Bufferizer) Write(buf []byte) (int, error) {
	b.Buff = append(b.Buff, buf...)
	return len(buf), nil
}

func (b *Bufferizer) Get() string {
	return string(b.Buff)
}
