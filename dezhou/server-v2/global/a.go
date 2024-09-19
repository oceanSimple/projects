package global

var (
	Players = newPlayerList()
	Rooms   = newRoomList()
)

func init() {
	initRoomList()

	initSomeRoom()
}
