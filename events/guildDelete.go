package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/pollen5/taiga/constants"
	"github.com/sapphire-cord/sapphire"
)

func GuildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	// If this event got triggered because the guild became unavailable
	// Add it to our cache so guildCreate can be aware when it becomes available and not treat it as a new guild.
	if g.Unavailable {
		guildsCache[g.ID] = true
		return
	}

	// Set the new presence.
	s.UpdateStatus(0, fmt.Sprintf("t!help | %d Servers!", len(s.State.Guilds)))

	// Unfortunately all the API gives us is just an ID and wether it is unavailable
	// And discordgo already cleared it from the cache by this point.
	// So best we can show is the ID and the created time.
	created, _ := discordgo.SnowflakeTimestamp(g.ID)

	s.ChannelMessageSendEmbed(constants.GuildLogsChannelID, sapphire.NewEmbed().
		SetTitle(fmt.Sprintf("%s has left a server.", s.State.User.Username)).
		SetDescription(g.ID).
		SetColor(0xFF0000).
		AddField("Created At", fmt.Sprintf("%s (%s)", created.Format("2 January 2006"), humanize.Time(created))).
		Build())
}
