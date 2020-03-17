package responses

// Responses used when checking level.

import (
	"github.com/pollen5/taiga/utils"
	"strconv"
)

var LevelMessages = []string{
	"You're only level **{{level}}**, Not bad...",
	"Pfft, level **{{level}}**?",
	"**{{user}}** your level? It's **{{level}}**, baka",
	"**{{user}}**, again with this? Whatever, you're level **{{level}}**.",
	"**{{user}}-san** here you go, it's **{{level}}**.",
	"You're level **{{level}}**, maybe if you were a wee bit more active it'd be higher.",
	"Only a level **{{level}}**? Even I'm better than that.",
	"Level **{{level}}**? What, where you hoping to be higher than **{{level}}**?",
	"Only level **{{level}}**? Pfft! I've seen much better than that.",
}

var OtherLevelMessages = []string{
	"Look just ask **{{user}}** next time, but I will forgive you. Here, this is their level: **{{level}}**.",
	"Why do you want to view **{{user}}**'s level? Anyway they are level **{{level}}**.",
}

var LevelUpMessages = []string{
	"You've only become level **{{level}}**, peasant!",
	"YATTA! You leveled up, **{{user}}-san**! You're now level **{{level}}**, n-not that it means a-anything!",
	"E-eh? You leveled up? Well just because you're level **{{level}}** now, doesn't mean I'll like you more **{{user}}**... baka.",
	"Hmph, you leveled up to **{{level}}**, **{{user}}-san**. W-well, it makes no difference to me!",
	"Hyaa~! You've leveled up **{{user}}-san**, you're now level **{{level}}**.",
	"**{{user}}**, you baka. Don't leave me behind now that you're level **{{level}}**.",
	"**{{user}}-san**, don't forget me now just because you're level **{{level}}**",
	"Pfft, you've only just become level **{{level}}**, hurry up **{{user}}**.",
	"Ah nuts, **{{user}}** is getting smart, now with an I.Q. of **{{level}}**",
	"N-nani? You've reached level **{{level}}** so quickly...",
	"You're only level **{{level}}**? You disgust me!",
	"How can you be level **{{level}}** y-you baka!",
	"B-baka! You're at level **{{level}}**, now go be free! D-dummy...",
	"You're definitely level **{{level}}** you baka!",
}

// Helpers to quickly roll.

func Level(username string, level int) string {
	return utils.NewString(utils.Roll(LevelMessages)).
		Replace("{{user}}", username).
		Replace("{{level}}", strconv.Itoa(level)).
		Build()
}

func OtherLevel(username string, level int) string {
	return utils.NewString(utils.Roll(OtherLevelMessages)).
		Replace("{{user}}", username).
		Replace("{{level}}", strconv.Itoa(level)).
		Build()
}

func LevelUp(username string, level int) string {
	return utils.NewString(utils.Roll(LevelUpMessages)).
		Replace("{{user}}", username).
		Replace("{{level}}", strconv.Itoa(level)).
		Build()
}
