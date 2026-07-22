package shell

import (
	"os/exec"

	"github.com/renatopp/go-x/fmtx"
)

type Result struct {
	output string
	err    error
	code   int
}

func NewResult(err error, combinedOutput, errorOutput string) Result {
	code := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		code = exitErr.ExitCode()
	}

	if err != nil {
		err = fmtx.Error("%s", errorOutput)
	}

	return Result{
		output: combinedOutput,
		err:    err,
		code:   code,
	}
}

func (r Result) Output() string {
	return r.output
}

func (r Result) Error() error {
	return r.err
}

func (r Result) ErrorWith(msg string) error {
	if r.err == nil {
		return nil
	}

	return fmtx.Error("%s: %s", msg, r.err.Error())
}

func (r Result) Code() int {
	return r.code
}

func (r Result) IsSuccess() bool {
	return r.err == nil
}

func (r Result) IsFailure() bool {
	return r.err != nil
}

func (r Result) IsCode(code int) bool {
	return r.code == code
}

func (r Result) String() string {
	return r.Output()
}

// func (r Result) Unwrap() Result {
// 	return r.UnwrapWith("")
// }

// func (r Result) UnwrapWith(msg string) Result {
// 	if r.err == nil {
// 		return r
// 	}

// 	if msg != "" {
// 		.Fatal(msg, "erro", r.Error().Error())
// 	} else {
// 		logger.Fatal(r.Error().Error())
// 	}
// 	return r
// }
