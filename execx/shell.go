package shell

import (
	"context"
	"fmt"
	"io"
	"maps"
	"os"
	"os/exec"
	"strings"
	"time"
)

var silent = false

func SetSilent(s bool) {
	silent = s
}

type Shell struct {
	command        string
	args           []string
	dir            string
	envs           map[string]string
	timeout        time.Duration
	stdin          io.Reader
	stdout         io.Writer
	stderr         io.Writer
	spinnerTitleFn func(string) string
}

// NewStd creates a new Shell instance with stdin/stdout/stderr set up to
// os.Stdin, os.Stdout, and os.Stderr, respectively.
func New(command string, args ...string) *Shell {
	return &Shell{
		command: command,
		args:    args,
		dir:     "",
		envs:    make(map[string]string),
		timeout: 0,
		stdin:   os.Stdin,
		stdout:  os.Stdout,
		stderr:  os.Stderr,
	}
}

// WithDir sets the working directory for the shell command.
func (s *Shell) WithDir(dir string) *Shell {
	s.dir = dir
	return s
}

// WithEnv sets an environment variable for the shell command.
func (s *Shell) WithEnv(key, value string) *Shell {
	s.envs[key] = value
	return s
}

// WithEnvs sets multiple environment variables for the shell command.
func (s *Shell) WithEnvs(envs map[string]string) *Shell {
	maps.Copy(s.envs, envs)
	return s
}

// WithTimeout sets a timeout for the shell command. If the command does not
// complete within the specified duration, it will be killed and an error will
// be returned.
func (s *Shell) WithTimeout(timeout time.Duration) *Shell {
	s.timeout = timeout
	return s
}

// WithStdin sets the standard input for the shell command.
func (s *Shell) WithStdin(stdin io.Reader) *Shell {
	s.stdin = stdin
	return s
}

// WithStdout sets the standard output for the shell command.
func (s *Shell) WithStdout(stdout io.Writer) *Shell {
	s.stdout = stdout
	return s
}

// WithStderr sets the standard error for the shell command.
func (s *Shell) WithStderr(stderr io.Writer) *Shell {
	s.stderr = stderr
	return s
}

// Cmd creates an exec.Cmd instance based on the Shell configuration.
func (s *Shell) Cmd() *exec.Cmd {
	cmd := exec.Command(s.command, s.args...)
	cmd.Dir = s.dir
	for k, v := range s.envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	if s.timeout > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
		defer cancel()
		cmd = exec.CommandContext(ctx, s.command, s.args...)
	}
	if s.stdin != nil {
		cmd.Stdin = s.stdin
	}
	if s.stdout != nil {
		cmd.Stdout = s.stdout
	}
	if s.stderr != nil {
		cmd.Stderr = s.stderr
	}
	return cmd
}

// Run executes the shell command and returns its combined output and any
// error that occurred during execution. If stdout or stderr is set, it
// will return an empty string for the output and the error from cmd.Run()
// instead.
func (s *Shell) Run() Result {
	cmd := s.Cmd()

	errs := strings.Builder{}
	combined := strings.Builder{}

	stdout := asWriter(func(str string) {
		if s.stdout != nil {
			s.stdout.Write([]byte(str))
		}
		combined.WriteString(str)
	})

	stderr := asWriter(func(str string) {
		if s.stderr != nil {
			s.stderr.Write([]byte(str))
		}
		combined.WriteString(str)
		errs.WriteString(str)
	})

	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return NewResult(err, combined.String(), errs.String())
}

type writer struct {
	fn func(s string)
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.fn(string(p))
	return len(p), nil
}

func asWriter(fn func(s string)) io.Writer {
	return &writer{fn: fn}
}
