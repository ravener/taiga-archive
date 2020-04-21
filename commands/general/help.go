package general

import (
	"github.com/sapphire-cord/sapphire"
	"github.com/bwmarrin/discordgo"
	"strings"
	"fmt"
)

// Shows a list of all commands.
// Usage: [command:string]
// Aliases: h, cmds, commands
func Help(ctx *sapphire.CommandContext) {
	if ctx.HasArgs() { // User passed an argument, give help information on that command only.
		cmd := ctx.Bot.GetCommand(ctx.Args[0].AsString())
		if cmd == nil {
			ctx.Reply("Unknown Command.")
			return
		}
		var aliases string = "None"

		if len(cmd.Aliases) > 0 {
			aliases = strings.Join(cmd.Aliases, ", ")
		}

		ctx.BuildEmbed(sapphire.NewEmbed().
			SetDescription("**Note:** Aisaka Taiga is going through a major rewrite which will get a name change and more features, the current one won't get developed further so please excuse any bugs and be patient for the new version. In the meantime [join our discord server](https://discord.gg/mDkMbEh) for updates.\n\n" + fmt.Sprintf("**Name:** %s\n**Description:** %s\n**Category:** %s\n**Aliases:** %s\n**Usage:** %s",
				cmd.Name,
				cmd.Description,
				cmd.Category,
				aliases,
				fmt.Sprintf("%s%s %s", ctx.Prefix, cmd.Name, sapphire.HumanizeUsage(cmd.UsageString)),
			)).SetColor(ctx.Bot.Color).SetTitle("Command Help"))
		return
	}
	// Send all commands.

	categories := make(map[string][]string)
	for _, v := range ctx.Bot.Commands {
		_, ok := categories[v.Category]
		if !ok {
			categories[v.Category] = []string{}
		}
		if !v.OwnerOnly || ctx.Author.ID == ctx.Bot.OwnerID {
			categories[v.Category] = append(categories[v.Category], v.Name)
		}
	}

	// Filter out empty categories, e.g a whole category is full of owner commands and the commands are filtered
	// because the user isn't the owner, so avoid even mentioning that empty category.
	for k, v := range categories {
		if len(v) == 0 {
			delete(categories, k)
		}
	}

	var embed = &discordgo.MessageEmbed{
		Title:  "Commands",
		Color:  ctx.Bot.Color,
		Footer: &discordgo.MessageEmbedFooter{Text: "For more info on a command use: " + ctx.Prefix + "help <command>"},
		Author: &discordgo.MessageEmbedAuthor{IconURL: ctx.Author.AvatarURL("256"), Name: ctx.Author.Username},
		Description: "**Note:** Aisaka Taiga is going through a major rewrite which will get a name change and more features, the current one won't get developed further so please excuse any bugs and be patient for the new version. In the meantime [join our discord server](https://discord.gg/mDkMbEh) for updates.",
	}

	for cat, cmds := range categories {
		var field = &discordgo.MessageEmbedField{Name: cat, Value: ""}
		field.Value = strings.Join(cmds, ", ")
		field.Inline = true
		embed.Fields = append(embed.Fields, field)
	}
	ctx.ReplyEmbed(embed)
}
