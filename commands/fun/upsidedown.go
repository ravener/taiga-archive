package fun

import (
	"github.com/sapphire-cord/sapphire"
	"github.com/pollen5/taiga/utils"
)

/*
¿‾ ¡z ʎ x ʍ ʌ n ʇ s ɹ b d o u ɯ l ʞ ɾ ı ɥ b ɟ ǝ p ɔ q ɐ Z ʎ X M Λ ∩ ⊥ S ᴚ Ὁ Ԁ O N W ˥ ʞ ſ I H ƃ Ⅎ Ǝ ᗡ Ͻ q ∀
*/

// TODO

// Turns your text into upside down.
// Aliases: upside, updown
// Usage: <text:string...>
func UpsideDown(ctx *sapphire.CommandContext) {
	ctx.Reply(utils.NewString(ctx.JoinedArgs()).
	  Replace("A", "∀").
	  Build())
}
