package global

import (
	"server-v2/model"
	"sync"
)

type RoomListClass struct {
	// 房间列表
	RoomList []*model.Room
	// 读写锁
	rwMutex *sync.RWMutex
	// 最小房间号
	MinRoomNumber int
	// 最大房间号
	MaxRoomNumber int
}

func newRoomList() *RoomListClass {
	return &RoomListClass{
		RoomList:      make([]*model.Room, 0),
		rwMutex:       &sync.RWMutex{},
		MinRoomNumber: 1,
		MaxRoomNumber: 100,
	}
}

func initRoomList() {
	for i := 0; i < 100; i++ {
		Rooms.RoomList = append(Rooms.RoomList, model.NewRoom(i+1))
	}
}

func (r *RoomListClass) RoomNumberIsValid(number int) bool {
	return number >= r.MinRoomNumber && number <= r.MaxRoomNumber
}

// 因为房间号是从1开始的, 所以要减1
func (r *RoomListClass) GetRoomFromRoomList(roomNumber int) *model.Room {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()
	if !r.RoomNumberIsValid(roomNumber) {
		return nil
	}
	return r.RoomList[roomNumber-1]
}
