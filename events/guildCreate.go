package events

import (
  "github.com/bwmarrin/discordgo"
  "github.com/pollen5/taiga/constants"
  "github.com/sapphire-cord/sapphire"
  "github.com/dustin/go-humanize"
  "strconv"
  "fmt"
)

func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
  if _, ok := guildsCache[g.ID]; ok {
    // An unavailable guild became available, delete it from our cache and continue.
    delete(guildsCache, g.ID)

    // Let us know when all guilds are loaded.
    if len(guildsCache) == 0 {
      logger.Infof("Done loading all guilds (count: %d)", len(s.State.Guilds))
    }

    return
  }

  // A new guild joined.
  owner, err := s.User(g.OwnerID)
  if err != nil { return }

  created, _ := discordgo.SnowflakeTimestamp(g.ID)

  s.ChannelMessageSendEmbed(constants.GuildLogsChannelID, sapphire.NewEmbed().
    SetTitle(fmt.Sprintf("%s has joined a new server!", s.State.User.Username)).
    SetDescription(g.Name).
    SetThumbnail(g.IconURL()).
    SetColor(0xDFAC7C).
    AddField("Owner", owner.String()).
    AddField("Member Count", strconv.Itoa(g.MemberCount)).
    AddField("Created At", fmt.Sprintf("%s (%s)", created.Format("2 January 2006"), humanize.Time(created))).
    SetFooter(g.ID).
    Build())
}
