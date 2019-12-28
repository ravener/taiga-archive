package events

import (
  "github.com/bwmarrin/discordgo"
  "fmt"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
  logger.Infof("Logged in as %s (%s)", r.User.String(), r.User.ID)
  s.UpdateStatus(0, fmt.Sprintf("t!help | %d Servers!", len(s.State.Guilds)))

  // Cache the unavailable guilds.
  for _, g := range r.Guilds {
    if g.Unavailable {
      guildsCache[g.ID] = true
    }
  }
}
