package social

import (
	"github.com/pollen5/taiga/constants/responses"
	"github.com/pollen5/taiga/db"
	"github.com/sapphire-cord/sapphire"
)

// View your balance or someone else's.
// Usage: [@@member]
// Aliases: bal, points
// guild only
func Balance(ctx *sapphire.CommandContext) {
	if ctx.HasArgs() {
		member := ctx.Arg(0).AsMember()
		name := member.User.Username

		if member.Nick != "" {
			name = member.Nick
		}

		ctx.Reply(responses.OtherBalance(name, db.GetBalance(member.User.ID, ctx.Guild.ID)))
		return
	}

	name := ctx.Author.Username
	member := ctx.Member(ctx.Author.ID)

	if member != nil && member.Nick != "" {
		name = member.Nick
	}

	ctx.Reply(responses.Balance(name, db.GetBalance(ctx.Author.ID, ctx.Guild.ID)))
}
