package model

import "time"

type SearchString struct {
	Block      string `json:"range"`
	SearchInfo string `json:"query"`
}

type ChatSearch struct { //只用于搜索
	RoomID        int      `gorm:"column:cid" json:"roomID"`
	RoomName      string   `gorm:"column:cname" json:"roomName"`
	RoomIntro     string   `gorm:"column:cremark" json:"roomIntro"`
	RoomOwnerID   int      `gorm:"column:uid" json:"roomOwnerID"`
	RoomOwnerName string   `gorm:"-" json:"roomOwnerName"` //房主名字还得去检索
	RoomOwnerIMg  string   `gorm:"-" json:"roomOwnerImg"`  //图片
	MemberList    []Member `gorm:"-" json:"memberList"`
}

type Room struct {
	Cid         int       `gorm:"column:cid;primaryKey;autoIncrement"`         //房间ID自增长不用传
	Name        string    `gorm:"column:cname;type:varchar(45);not null"`      //房名 要给
	Bid         int       `gorm:"column:bid;not null"`                         //版块ID 要给
	Uid         int       `gorm:"column:uid;not null"`                         //房主ID 不用给
	Cnumber     int       `gorm:"column:cnumber;not null"`                     //人数 要给
	Cstate      int       `gorm:"column:cstate;not null"`                      //状态默认为1  不用给
	CYueDate    string    `gorm:"column:c_yue_date;type:varchar(12);not null"` //约的时间 要给
	CCreateTime time.Time `gorm:"column:c_create_time;type:datetime;not null"` //创建时间 不用给
	Cremark     string    `gorm:"column:cremark;type:text"`                    //房间简介 要给
}

type Member struct {
	MemberID  int    `gorm:"column:uid" json:"memberID"`
	MemberImg string `gorm:"-" json:"memberImg"`
}
