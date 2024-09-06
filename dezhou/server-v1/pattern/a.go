package pattern

import (
	"server-v1/card"
	"server-v1/tool"
)

var patternJudge = make([]CardType, 0, 9)

func init() {
	patternJudge = append(patternJudge, NewStraightFlush())
	patternJudge = append(patternJudge, NewFourOfKind())
	patternJudge = append(patternJudge, NewFullHouse())
	patternJudge = append(patternJudge, NewFlush())
	patternJudge = append(patternJudge, NewStraight())
	patternJudge = append(patternJudge, NewThreeOfKind())
	patternJudge = append(patternJudge, NewTwoPairs())
	patternJudge = append(patternJudge, NewOnePair())
	patternJudge = append(patternJudge, NewSingleCard())
}

func Judge(c []*card.Card) (weight int, arr []*card.Card) {
	for i := 0; i < len(patternJudge); i++ {
		tool.SortCardsByNumber(c)
		flag, cards := patternJudge[i].Judge(c)
		if flag {
			return patternJudge[i].GetWeight(), cards
		}
	}
	return -1, nil
}
