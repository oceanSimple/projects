package cardSample

import "server-v1/card"

// 皇家同花顺：♥A、♥K、♥Q、♥J、♥10
func A_K_Q_J_10_6_4() []*card.Card {
	return []*card.Card{
		&card.Card{Number: card.Number.Ace, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.King, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Queen, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Ten, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Six, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Four, Color: card.Color.Hearts},
	}
}

// 同花顺：♥A、♥J、♥10、♥5、♥4、♥3、♥2
func A_J_10_5_4_3_2() []*card.Card {
	return []*card.Card{
		&card.Card{Number: card.Number.Ace, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Ten, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Five, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Four, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Three, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Two, Color: card.Color.Hearts},
	}
}

// 同花顺：♥10、♥9、♥8、♥7、♥6、♥5
func A_10_9_8_7_6_5() []*card.Card {
	return []*card.Card{
		&card.Card{Number: card.Number.Ace, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Ten, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Nine, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Eight, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Seven, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Six, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Five, Color: card.Color.Hearts},
	}
}

// 非同花： ♥A、♠J、♥10、♥5、♣4、♥3、♥2
func A_J_10_5_4_3_2_NoFlush() []*card.Card {
	return []*card.Card{
		&card.Card{Number: card.Number.Ace, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Spades},
		&card.Card{Number: card.Number.Ten, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Five, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Four, Color: card.Color.Clubs},
		&card.Card{Number: card.Number.Three, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Two, Color: card.Color.Hearts},
	}
}

// 四条： ♥K, ♥Q, ♥J, ♠J, ♣J, ♦J，♥5
func K_Q_J_J_J_J_5() []*card.Card {
	return []*card.Card{
		&card.Card{Number: card.Number.King, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Queen, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Spades},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Clubs},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Diamonds},
		&card.Card{Number: card.Number.Five, Color: card.Color.Hearts},
	}
}

// 葫芦：♥K, ♥Q, ♥J, ♠J, ♣J, ♦5，♥5
func K_Q_J_J_J_5_5() []*card.Card {
	return []*card.Card{
		&card.Card{Number: card.Number.King, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Queen, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Spades},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Clubs},
		&card.Card{Number: card.Number.Five, Color: card.Color.Diamonds},
		&card.Card{Number: card.Number.Five, Color: card.Color.Hearts},
	}
}

// 特殊葫芦：♥K,  ♥J, ♠J, ♣J, ♦5, ♥5, ♠5
func K_J_J_J_5_5_5() []*card.Card {
	return []*card.Card{
		&card.Card{Number: card.Number.King, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Spades},
		&card.Card{Number: card.Number.Jack, Color: card.Color.Clubs},
		&card.Card{Number: card.Number.Five, Color: card.Color.Diamonds},
		&card.Card{Number: card.Number.Five, Color: card.Color.Hearts},
		&card.Card{Number: card.Number.Five, Color: card.Color.Spades},
	}
}
