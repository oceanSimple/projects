package patternTool

import (
	"server-v2/card"
	"server-v2/tool"
)

var HasStraight PatternToolInterface = func(cards []card.Card) (bool, [][]card.Card) {
	var result [][]card.Card

	commonStraight := getStraight(cards)
	specialStraight := getSpecialStraight(cards)

	result = append(result, commonStraight...)
	result = append(result, specialStraight...)

	return len(result) > 0, result
}

// 查找顺子
func getStraight(cards []card.Card) [][]card.Card {
	var result [][]card.Card

	initialArray := getInitialArray(cards)

	for i := 0; i < len(initialArray); i++ {
		result = append(result, queueDfs(initialArray[i])...)
	}

	return result
}

// 查找特殊的顺子：5-4-3-2-A
func getSpecialStraight(cards []card.Card) [][]card.Card {
	// 暂时将A的number.Weight设置为1
	for i := 0; i < len(cards); i++ {
		if cards[i].Number.Weight == 14 {
			cards[i].Number.Weight = 1
		} else {
			break
		}
	}
	// 按照权重排序
	tool.SortCardsByNumberExtraWeight(cards)

	// 查找顺子
	// 为了避免与常规的顺子重复，我们从weight=5的牌开始查找
	var index = 0
	for ; index < len(cards); index++ {
		if cards[index].Number.Weight == 5 {
			break
		}
	}
	result := getStraight(cards[index:])
	// 恢复权重
	for i := len(cards) - 1; i >= 0; i-- {
		if cards[i].Number.Weight == 1 {
			cards[i].Number.Weight = 14
		} else {
			break
		}
	}
	// 重新排序
	tool.SortCardsByNumberWeight(cards)
	return result
}

// 举个例子：皇家同花顺：♠A-♥A-♠K-♥K-♠Q-♠J-♠10
// initialArray = [
// [[♠A], [♠K, ♥K], [♠Q], [♠J], [♠10]],
// [[♥A], [♠K, ♥K], [♠Q], [♠J], [♠10]],
// ]
func getInitialArray(cards []card.Card) [][][]card.Card {
	var initialArray [][][]card.Card

	// 五张为一组顺子，我们只需要遍历到倒数第五张牌即可
	for i := 0; i < len(cards)-4; i++ {
		var tempArray [][]card.Card
		tempArray = append(tempArray, []card.Card{cards[i]}) // 将第一张牌放入tempArray

		// 索引
		var index = i + 1
		// 标准值
		var standard = cards[i].Number.Weight

		// 依次寻找比当前牌小1的牌，直到找到5张牌
		for j := 4; j > 0; j-- {
			standard--
			nextIndex, cardArray := getTargetCardArray(cards, index, len(cards)-j, standard)
			if len(cardArray) == 0 {
				break
			}
			index = nextIndex
			tempArray = append(tempArray, cardArray)
		}

		// 判断是否找到了剩下的4张牌
		if len(tempArray) == 5 {
			initialArray = append(initialArray, tempArray)
		}
	}
	return initialArray
}

// 对于[[♠A], [♠K, ♥K], [♠Q], [♠J], [♠10]]
// 返回[ [♠A, ♠K, ♠Q, ♠J, ♠10], [♠A, ♥K, ♠Q, ♠J, ♠10] ]
func queueDfs(cards [][]card.Card) [][]card.Card {
	var result = make([][]card.Card, 0)
	// 将第一张牌放入result
	result = append(result, cards[0])

	// 从第二组牌开始遍历
	// result中已有的牌称为prefix，第二组牌中的每一张牌称为suffix
	// 将prefix和suffix组合成一组牌，放入result
	for i := 1; i < len(cards); i++ {
		// 用来存储组合后的牌
		var temp = make([][]card.Card, 0)
		// 分别递归suffix和prefix，并进行组合
		for _, suffix := range cards[i] {
			for _, prefix := range result {
				temp = append(temp, append(prefix, suffix))
			}
		}
		// 将temp赋值给result
		result = temp
	}
	return result
}

// 获取card.Number.Weight等于standard的牌
// param cards: 所有的牌
// param start: 开始索引
// param end: 结束索引
// param standard: 标准值
// return int: 结束时的索引
// return []card.Card: 符合条件的牌
func getTargetCardArray(cards []card.Card, start, end, standard int) (int, []card.Card) {
	var result []card.Card
	var i int
	for i = start; i <= end; i++ {
		// 三种情况
		// 1. 当前牌的权重等于标准值，将当前牌放入result，继续还有没有符合的牌
		// 2. 当前牌的权重小于标准值，直接break。因为我们是从大到小遍历的，所以后面的牌肯定不符合
		// 3. 当前牌的权重大于标准值，继续遍历
		if cards[i].Number.Weight == standard {
			result = append(result, cards[i])
		} else if cards[i].Number.Weight < standard {
			break
		} else {
			continue
		}
	}
	return i, result
}
