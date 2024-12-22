package model

type SearchString struct {
	Block      string `json:"range"`
	SearchInfo string `json:"query"`
}

type ChatInfo struct {
	RoomID        int      `gorm:"column:cid" json:"roomID"`
	RoomName      string   `gorm:"column:cname" json:"roomName"`
	RoomIntro     string   `gorm:"column:cremark" json:"roomIntro"`
	RoomOwnerID   int      `gorm:"column:uid" json:"roomOwnerID"`
	RoomOwnerName string   `gorm:"-" json:"roomOwnerName"` //房主名字还得去检索
	RoomOwnerIMg  string   `gorm:"-" json:"roomOwnerImg"`  //图片
	MemberList    []Member `gorm:"-" json:"memberList"`
}

type Member struct {
	MemberID  int    `gorm:"column:uid" json:"memberID"`
	MemberImg string `gorm:"-" json:"memberImg"`
}
