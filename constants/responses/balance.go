package responses

import (
	"github.com/pollen5/taiga/utils"
	"github.com/dustin/go-humanize"
)

var BalanceMessages = []string{
	"Pfft, a measly **¥{{amount}}**? My father gives me more as an allowance!",
	"Only **¥{{amount}}**? That's nothing compared to my pocket money!",
	"**{{user}}-san**, you have **¥{{amount}}**, but you don't need it to make me happy.",
	"**{{user}}**, again? Ugh, you have **¥{{amount}}** ... Jeez, maybe if you kept track of it you'd remember...",
}

var OtherBalanceMessages = []string{
	"Why do you want to know **{{user}}**'s balance? Whatever, it's **¥{{amount}}** anyway.",
	"N-nani? **{{user}}-san**'s balance? It's **¥{{amount}}**..",
	"Why do you care so much about **{{user}}**'s balance? N-not that I care, it's **¥{{amount}}**.",
	"Some particular reason you want to know **{{user}}**'s balance? It's **¥{{amount}}**.",
	"Hm, one second... It's **¥{{amount}}**, but why do you want to know **{{user}}**'s balance?",
	"You've aleady asked for **{{user}}**'s balance, gosh! Anyway, it's **¥{{amount}}**.",
}

func Balance(username string, amount int64) string {
	return utils.NewString(utils.Roll(BalanceMessages)).
		Replace("{{user}}", username).
		Replace("{{amount}}", humanize.Comma(amount)).
		Build()
}

func OtherBalance(username string, amount int64) string {
	return utils.NewString(utils.Roll(OtherBalanceMessages)).
		Replace("{{user}}", username).
		Replace("{{amount}}", humanize.Comma(amount)).
		Build()
}
