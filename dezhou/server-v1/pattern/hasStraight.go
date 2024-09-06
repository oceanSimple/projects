package pattern

import (
	"server-v1/card"
	"server-v1/tool"
)

func HasStraight(cards []*card.Card) (bool, [][]*card.Card) {
	if len(cards) != 7 {
		return false, nil
	}

	var result = make([][]*card.Card, 0)
	var middleArray = make([][][]*card.Card, 0)

	for i := 0; i < 3; i++ {
		var temp = make([][]*card.Card, 0)
		temp = append(temp, []*card.Card{cards[i]})
		var standard = cards[i].Number.Weight - 1
		var index = i + 1

		for j := 0; j < 4; j++ {
			_, cardArray := getCardNArray(cards, index, 3+j, standard)
			if len(cardArray) == 0 {
				break
			}
			temp = append(temp, cardArray)
			standard--
		}

		if len(temp) == 5 {
			middleArray = append(middleArray, temp)
		}
	}

	for _, arr := range middleArray {
		result = append(result, queueDfs(arr)...)
	}

	// get special straight
	specialFlag, specialStraight := specialStraight(cards)
	if specialFlag {
		result = append(result, specialStraight...)
	}

	if len(result) > 0 {
		return true, result
	}
	return false, nil
}

// 5-4-3-2-A
func specialStraight(cards []*card.Card) (bool, [][]*card.Card) {
	// sort cards by extra weight
	tool.SortCardsByNumberExtraWeight(cards)
	// recover the sort by weight
	defer tool.SortCardsByNumber(cards)

	var result = make([][]*card.Card, 0)
	var middleArray = make([][][]*card.Card, 0)

	for i := 0; i < 3; i++ {
		var temp = make([][]*card.Card, 0)
		temp = append(temp, []*card.Card{cards[i]})
		var standard = cards[i].Number.ExtraWeight - 1
		var index = i + 1

		for j := 0; j < 4; j++ {
			_, cardArray := func(mainCardArray []*card.Card, start, end, standard int) (int, []*card.Card) {
				var result = make([]*card.Card, 0)
				for i := start; i <= end; i++ {
					if mainCardArray[i].Number.ExtraWeight == standard {
						result = append(result, mainCardArray[i])
					} else if mainCardArray[i].Number.ExtraWeight < standard {
						break
					}
				}
				return start + len(result), result
			}(cards, index, 3+j, standard)

			if len(cardArray) == 0 {
				break
			}
			temp = append(temp, cardArray)
			standard--
		}

		if len(temp) == 5 {
			middleArray = append(middleArray, temp)
		}
	}

	for _, arr := range middleArray {
		result = append(result, queueDfs(arr)...)
	}

	// only return the array start with 5
	var temp = make([][]*card.Card, 0)
	for i := 0; i < len(result); i++ {
		if result[i][0].Number.ExtraWeight == 5 {
			temp = append(temp, result[i])
		}
	}

	if len(temp) > 0 {
		return true, temp
	}
	return false, nil
}

func getCardNArray(mainCardArray []*card.Card, start, end, standard int) (int, []*card.Card) {
	var result = make([]*card.Card, 0)
	for i := start; i <= end; i++ {
		if mainCardArray[i].Number.Weight == standard {
			result = append(result, mainCardArray[i])
		} else if mainCardArray[i].Number.Weight < standard {
			break
		}
	}
	return start + len(result), result
}

func queueDfs(cardArray [][]*card.Card) [][]*card.Card {
	l := make([][]*card.Card, 0, len(cardArray))
	l = append(l, cardArray[0])
	for i := 1; i < len(cardArray); i++ {
		var temp = make([][]*card.Card, 0)
		for _, suffix := range cardArray[i] {
			for _, prefix := range l {
				temp = append(temp, append(prefix, suffix))
			}
		}
		l = temp
	}
	return l
}
