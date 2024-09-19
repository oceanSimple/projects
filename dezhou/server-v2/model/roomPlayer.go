package model

type RoomPlayer struct {
	Account    string // 玩家的账号
	Chip       int    // 玩家的筹码
	IsPrepared int    // 玩家是否准备: 0-未准备, 1-已准备
	IsHost     int    // 玩家是否是房主: 0-不是, 1-是
}
