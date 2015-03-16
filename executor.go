package meechum

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
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
func (e *Executor) Do(p string, args []string) (string, error) {

	buffer := NewBufferizer()

	// Set buffers for this run
	stdout := io.MultiWriter(os.Stdout, buffer)
	stderr := io.MultiWriter(os.Stderr, buffer)

	// Setup work directory and command
	execPath := path.Clean(e.Workdir + "/" + p)
	cmd := &exec.Cmd{
		Dir:    execPath,
		Path:   e.Binary,
		Args:   args,
		Stdout: stdout,
		Stderr: stderr,
	}

	// Log and execute
	log.Infof("Exec %s %s %s", e.Binary, strings.Join(args, " "), execPath)

	err := cmd.Run()
	if err != nil {
		return buffer.Get(), err
	}
	return buffer.Get(), nil
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
