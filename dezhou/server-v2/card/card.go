package card

import (
	"fmt"
	"server-v2/output"
)

type Card struct {
	Color  ColorObj  // 颜色
	Number NumberObj // 数字
}

func (c Card) String() string {
	return fmt.Sprintf("{%s, %s}", c.Color.Picture, c.Number.Text)
}

// 根据花色对控制台输出进行着色
func (c Card) DyeString() string {
	if c.Color.Picture == "♠" || c.Color.Picture == "♣" {
		return fmt.Sprintf("{%s, %s}", output.DyeText(c.Color.Picture, output.Grey), c.Number.Text)
	} else {
		return fmt.Sprintf("{%s, %s}", output.DyeText(c.Color.Picture, output.Red), c.Number.Text)
	}
}

// 由于ColorObj和NumberObj只能从初始化的enum中获取
// 因此我们只需要比较牌的数字和花色是否相同即可
func (c Card) Equal(another Card) bool {
	return c.Number.Weight == another.Number.Weight &&
		c.Color.Text == another.Color.Text
}
