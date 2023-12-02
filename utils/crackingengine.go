package utils

import (
   // "fmt"
   "strings"
)

type CrackingEngine struct {
	Executable string
    Args []string
}

func (e *CrackingEngine) AddArg(Arg string) {
    e.Args = append(e.Args, Arg)
}

func (e *CrackingEngine) AsCmd(Arg string) string {
    for i, arg := range e.Args {
        if strings.ContainsAny(arg, " ?") {
			e.Args[i] = `"` + arg + `"`
		}
    }
    return e.Executable + " " + strings.Join(e.Args, " ")
}

