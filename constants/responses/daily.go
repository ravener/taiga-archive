package responses

import (
	"github.com/pollen5/taiga/utils"
)

var DailySuccessMessages = []string{
	"Yatta! You've got **{{amount}}**.",
	"Well done, you've redeemed your daily **{{amount}}**!",
	"Finally I thought you were never going to claim your **{{amount}}** today!",
	"Your dad gave you **{{amount}}**? I get more than that every hour!",
	"You have claimed your daily **{{amount}}**, ain't that dandy?",
	"N-nani?! You got **{{amount}}**, woah...",
	"You just got **{{amount}}**? Maybe buy me dinner some time **{{user}}-san**? :wink:",
	"Oh goody, you finally got your **{{amount}}**. It's about time, now get me some Pocky.",
}

var DailyFailureMessages = []string{
	"You cannot claim your daily reward yet, please try again in **{{time}}**",
	"Oh come on now, you know better than to ask ahead of time.. You can get your daily in **{{time}}**",
	"**{{user}}-san**, you already got your allowance today. You can get it in **{{time}}**",
	"This again? I told you to wait **{{time}}**",
	"Hey! Money is already tight around here. Ask in **{{time}}**",
	"You're gonna make m-me mad please wait **{{time}}** to claim again!",
}

func DailySuccess(user, time string) string {
	return utils.NewString(utils.Roll(DailySuccessMessages)).
		Replace("{{user}}", user).
		Replace("{{time}}", time).
		Build()
}

func DailyFailure(user, time string) string {
	return utils.NewString(utils.Roll(DailyFailureMessages)).
		Replace("{{user}}", user).
		Replace("{{time}}", time).
		Build()
}
