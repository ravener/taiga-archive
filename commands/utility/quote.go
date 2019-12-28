package utility

import (
  "github.com/sapphire-cord/sapphire"
  "fmt"
)

// Quote a message by ID.
// Usage: <message:string>
// guild only
func Quote(ctx *sapphire.CommandContext) {
  msg, err := ctx.Session.ChannelMessage(ctx.Channel.ID, ctx.Arg(0).AsString())

  if err != nil {
    ctx.Reply("Failed to fetch that message, is the ID correct?")
    return
  }

  ctx.BuildEmbed(sapphire.NewEmbed().
    SetTitle("Message Quote").
    SetURL(fmt.Sprintf("https://discordapp.com/channels/%s/%s/%s", ctx.Guild.ID, ctx.Channel.ID, msg.ID)).
    SetDescription(msg.Content).
    SetAuthor(msg.Author.Username, msg.Author.AvatarURL("256")).
    SetColor(0xDFAC7C).
    SetThumbnail(msg.Author.AvatarURL("2048")))
}
