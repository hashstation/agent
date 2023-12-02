package utils

import (
   // "fmt"
   "strings"
)

type Hashcat struct {
	Executable string
    Args []string
}

func (h *Hashcat) AddArg(Arg string) {
    h.Args = append(h.Args, Arg)
}

func (h *Hashcat) AsCmd(Arg string) string {
    for i, arg := range h.Args {
        if strings.ContainsAny(arg, " ?") {
			h.Args[i] = `"` + arg + `"`
		}
    }
    return h.Executable + " " + strings.Join(h.Args, " ")
}

