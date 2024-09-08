package patternTool

import "server-v2/card"

type PatternToolInterface func(cards []card.Card) (bool, [][]card.Card)
