package cs

type Waiter struct {
	Uid      int64
	Nickname string
	Avatar   string
}

type Data struct {
	Sign string `json:"sign"`
}

type RoomRequest struct {
	Name string `json:"name"`
}
