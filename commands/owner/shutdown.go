package owner

import (
	"github.com/pollen5/taiga/db"
	"github.com/sapphire-cord/sapphire"
	"os"
)

// Shuts down the bot.
// Aliases: reboot
func OwnerShutdown(ctx *sapphire.CommandContext) {
	ctx.Reply("Shutting down...")

	// Cleanly close the session and the database.
	ctx.Session.Close()
	db.Close()

	os.Exit(0)
}
