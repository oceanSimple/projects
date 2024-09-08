package pattern

import "server-v2/card"

var patternJudge = make([]Pattern, 0, 9)

type Pattern interface {
	// 判断传入的牌是否符合当前牌型
	// params： cards 传入的牌（规定为7张，2张手牌+五张公共牌）
	// return： bool 是否符合当前牌型
	// return： []card.Card 符合当前牌型的牌，用于相同牌型时Compare函数进行比大小（每种牌型返回的牌数不同）
	Judge(cards []card.Card) (bool, []card.Card)

	// 比较两个符合当前牌型的牌的大小
	// params： cards1 第一个符合当前牌型的牌
	// params： cards2 第二个符合当前牌型的牌
	// return： int 1：cards1大，-1：cards2大，0：相等
	// Compare(cards1 []card.Card, cards2 []card.Card) int

	// 获取当前牌型的权重
	// return： int 权重
	GetWeight() int
}

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

func Judge(c []card.Card) (weight int, arr []card.Card) {
	for i := 0; i < len(patternJudge); i++ {
		flag, cards := patternJudge[i].Judge(c)
		if flag {
			return patternJudge[i].GetWeight(), cards
		}
	}
	return -1, nil
}
