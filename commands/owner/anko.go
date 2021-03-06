package owner

import (
	"bytes"
	"fmt"
	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
	"github.com/sapphire-cord/sapphire"
)

// Evaluates arbitrary Anko
// Usage: <code:string...>
func OwnerAnko(ctx *sapphire.CommandContext) {
	e := env.NewEnv()
	stdout := &bytes.Buffer{}

	// define also returns an error but it really shouldn't happen in our case so we ignore it.

	e.Define("println", func(args ...interface{}) {
		stdout.WriteString(fmt.Sprintln(args...))
	})

	e.Define("ctx", ctx)
	e.Define("bot", ctx.Bot)
	e.Define("session", ctx.Session)

	value, err := vm.Execute(e, nil, ctx.JoinedArgs())
	if err != nil {
		ctx.CodeBlock("go", "%v\n", err)
		return
	}

	ctx.CodeBlock("go", "%s%+v", stdout.String(), value)
}
