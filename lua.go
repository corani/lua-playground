package main

import (
	"bytes"
	"context"

	lua "github.com/yuin/gopher-lua"
)

func executeLua(ctx context.Context, input string) (string, error) {
	state := lua.NewState()
	defer state.Close()

	state.SetContext(ctx)

	var buf bytes.Buffer

	// NOTE(daniel): redirect `print` output to buffer instead of writing it to stdout.
	// Adapted from: https://github.com/xyproto/algernon/blob/34d3806bfa890e7f6296cc400e8fd51951cdc926/basic.go#L67
	state.SetGlobal("print", state.NewFunction(func(L *lua.LState) int {
		top := L.GetTop()
		for i := 1; i <= top; i++ {
			buf.WriteString(L.Get(i).String())
			if i != top {
				buf.WriteString("\t")
			}
		}
		// Final newline
		buf.WriteString("\n")

		return 0 // number of results
	}))

	if err := state.DoString(input); err != nil {
		return "", err
	}

	return buf.String(), nil
}
