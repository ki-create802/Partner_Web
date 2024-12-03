package model

type SearchString struct {
	SearchInfo string `json:"query"`
}

type ChatInfo struct {
	RoomID        int      `gorm:"cid" json:"roomID"`
	RoomName      string   `gorm:"column:cname" json:"roomName"`
	RoomIntro     string   `gorm:"column:cremark" json:"roomIntro"`
	RoomOwnerID   int      `gorm:"column:uid" json:"roomOwnerID"`
	RoomOwnerName string   `gorm:"-" json:"roomOwnerName"` //房主名字还得去检索
	MemberList    []Member `gorm:"-" json:"memberList"`
}

type Member struct {
	MemberID  int    `gorm:"column:uid" json:"memberID"`
	MemberImg string `gorm:"-" json:"memberImg"`
}
