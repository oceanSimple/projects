package pattern

import (
	"server-v1/card"
	"server-v1/tool"
)

func HasFlush(cards []*card.Card) (bool, [][]*card.Card) {
	if len(cards) != 7 {
		return false, nil
	}

	var result = make([][]*card.Card, 0)
	// Sort cards by color
	tool.SortCardsByColor(cards)
	// defer sort cards by number
	// defer tool.SortCardsByNumber(cards)

	for i := 0; i < 3; i++ {
		for j := i + 1; j <= i+4; j++ {
			if cards[j].Color != cards[j-1].Color {
				break
			}
			if j == i+4 {
				// Sort cards by number and then append to result
				temp := cards[i : i+5]
				tool.SortCardsByNumber(temp)
				result = append(result, temp)
			}
		}
	}
	if len(result) > 0 {
		return true, result
	} else {
		return false, nil
	}
}

func HasFourSame(cards []*card.Card) (bool, [][]*card.Card) {
	if len(cards) != 7 {
		return false, nil
	}

	var result = make([][]*card.Card, 0)

	for i := 0; i < 4; i++ {
		if cards[i].Number.Weight == cards[i+3].Number.Weight {
			result = append(result, cards[i:i+4])
			break // only 7 cards, so only one four of kind
		}
	}
	if len(result) > 0 {
		return true, result
	} else {
		return false, nil
	}
}

func HasThreeSame(cards []*card.Card) (bool, [][]*card.Card) {
	if len(cards) != 7 {
		return false, nil
	}

	var result = make([][]*card.Card, 0)

	for i := 0; i < 5; i++ {
		if cards[i].Number.Weight == cards[i+2].Number.Weight {
			result = append(result, cards[i:i+3])
		}
	}
	if len(result) > 0 {
		// Deduplication
		result = Deduplication(result)
		return true, result
	} else {
		return false, nil
	}
}

func HasTwoSame(cards []*card.Card) (bool, [][]*card.Card) {
	if len(cards) != 7 {
		return false, nil
	}

	var result = make([][]*card.Card, 0)

	for i := 0; i < 6; i++ {
		if cards[i].Number.Weight == cards[i+1].Number.Weight {
			result = append(result, cards[i:i+2])
		}
	}
	if len(result) > 0 {
		// Deduplication
		result = Deduplication(result)
		return true, result
	} else {
		return false, nil
	}
}

// Remove duplicate cards
func Deduplication(cards [][]*card.Card) [][]*card.Card {
	var mapCards = make(map[int]bool)
	var result = make([][]*card.Card, 0)
	for i := 0; i < len(cards); i++ {
		var key = cards[i][0].Number.Weight
		if _, ok := mapCards[key]; !ok {
			mapCards[key] = true
			result = append(result, cards[i])
		}
	}
	return result
}
