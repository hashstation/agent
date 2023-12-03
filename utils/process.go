package utils

import (
	"strings"
)

type Process struct {
	Executable string
	Args       []string
}

func NewProcess(executable string) *Process {
	if executable == "" {
		return nil
	}
	return &Process{
		Executable: executable,
	}
}

func (e *Process) AddArg(Arg string) {
	e.Args = append(e.Args, Arg)
}

func (e *Process) AsCmd(Arg string) string {
	for i, arg := range e.Args {
		if strings.ContainsAny(arg, " ?") {
			e.Args[i] = `"` + arg + `"`
		}
	}
	return e.Executable + " " + strings.Join(e.Args, " ")
}
