package types

type Rules struct {
	Ants  int
	Rooms []Room
	Links []Link
	Start Room
	End   Room
}

type Room struct {
	Room     string
	Location [2]int
}

type Link struct {
	RoomA string
	RoomB string
}

type AntStatus struct {
	Ant              int
	Tunnel           int
	CurrentRoom      string
	CurrentRoomIndex int
	Finished         bool
}
