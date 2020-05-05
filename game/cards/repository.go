package cards

import (
	"duel-masters/game/cards/dm01"
	"duel-masters/game/match"
)

// Cards is a map with all the card id's in the game and corresponding CardConstructor
var Cards = map[string]match.CardConstructor{
	"57eeb3c3-2561-4841-a381-2e50d17533d1": dm01.AquaHulcus,
}
