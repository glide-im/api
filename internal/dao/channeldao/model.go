package channeldao

type ChannelModel struct {
	ChanId    string `gorm:"primaryKey" json:"chan_id"`
	Name      string
	Avatar    string
	Muted     bool
	Type      int
	ReadOnly  bool `json:"read_only"`
	Access    int
	UpdatedAt int64 `json:"updated_at"`
	CreatedAt int64 `json:"created_at"`
}

type ChannelMemberModel struct {
	MemberId  string `gorm:"primaryKey" json:"member_id"`
	ChanId    string `json:"chan_id"`
	Uid       string
	Muted     bool
	Type      int
	Perm      int64
	UpdatedAt int64 `json:"updated_at"`
	CreatedAt int64 `json:"created_at"`
}
