package events

import (
  "github.com/bwmarrin/discordgo"
  "github.com/sapphire-cord/sapphire"
  "github.com/pollen5/taiga/constants"
  "github.com/dustin/go-humanize"
  "strconv"
  "fmt"
)

func GuildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
  // If this event got triggered because the guild became unavailable
  // Add it to our cache so guildCreate can be aware when it becomes available and not treat it as a new guild.
  if g.Unavailable {
    guildsCache[g.ID] = true
    return
  }

  owner, err := s.User(g.OwnerID)
  if err != nil { return }

  created, _ := discordgo.SnowflakeTimestamp(g.ID)

  s.ChannelMessageSendEmbed(constants.GuildLogsChannelID, sapphire.NewEmbed().
    SetTitle(fmt.Sprintf("%s has left a server.", s.State.User.Username)).
    SetDescription(g.Name).
    SetThumbnail(g.IconURL()).
    SetColor(0xFF0000).
    AddField("Owner", owner.String()).
    AddField("Member Count", strconv.Itoa(g.MemberCount)).
    AddField("Created At", fmt.Sprintf("%s (%s)", created.Format("2 January 2006"), humanize.Time(created))).
    SetFooter(g.ID).
    Build())
}
