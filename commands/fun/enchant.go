package fun

import (
	"github.com/sapphire-cord/sapphire"
	"github.com/pollen5/taiga/utils"
)

// Convert your text into minecraft enchantment table language
// Aliases: enchantify
// Usage: <text:string...>
func Enchant(ctx *sapphire.CommandContext) {
	ctx.Reply(utils.NewString(ctx.JoinedArgs()).
	  Replace("a", "ᔑ").
		Replace("b", "ʖ").
    Replace("c", "ᓵ").
		Replace("d", "↸").
		Replace("e", "ᒷ").
    Replace("f", "⎓").
		Replace("g", "⊣").
		Replace("h", "⍑").
    Replace("i", "╎").
    Replace("j", "⋮").
		Replace("k", "ꖌ").
    Replace("l", "ꖎ").
    Replace("m", "ᒲ").
    Replace("n", "リ").
    Replace("o", "𝙹").
    Replace("p", "!¡").
		Replace("q", "ᑑ").
    Replace("r", "∷").
    Replace("s", "ᓭ").
    Replace("t", "ℸ ̣").
    Replace("u", "⚍").
    Replace("v", "⍊").
    Replace("w", "∴").
    Replace("x", "·/").
    Replace("y", "||").
    Replace("z", "⨅").
		Build())
}
