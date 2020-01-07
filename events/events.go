package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pollen5/minori"
)

var logger = minori.GetLogger("Events")

// Used to lookup if a guild is unavailable.
// It is used in guildCreate to know if an unavaiable guild has become available or it's a new guild joined.
// The value has no meaning, we simply use a map for it's constant O(1) lookup time.
var guildsCache = make(map[string]bool)

func Init(s *discordgo.Session) {
	s.AddHandler(Ready)
	s.AddHandler(GuildCreate)
	s.AddHandler(GuildDelete)
}
