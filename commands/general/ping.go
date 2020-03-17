package general

import (
	"github.com/pollen5/taiga/constants/responses"
	"github.com/sapphire-cord/sapphire"
)

func Ping(ctx *sapphire.CommandContext) {
	msg, err := ctx.ReplyLocale("COMMAND_PING")
  // Should never happen but if it did, avoid panics.
  if err != nil {
		return
  }

	usertime, err := ctx.Message.Timestamp.Parse()
  if err != nil {
    return
  }

  bottime, err := msg.Timestamp.Parse()
  if err != nil {
    return
  }

	ctx.Reply(responses.Ping(ctx.Author.Username, bottime.Sub(usertime).Milliseconds()))
}
