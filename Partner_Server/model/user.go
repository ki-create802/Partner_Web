package model

type User struct {
	UID     uint   `gorm:"column:uid;not null;primaryKey"`
	UName   string `gorm:"column:uname;type:varchar(45);not null"`
	UKey    string `gorm:"column:ukey;type:varchar(45);not null" json:"Password"`
	UEmail  string `gorm:"column:uemail;type:varchar(45);not null" json:"Email"`
	URemark string `gorm:"column:uremark;type:text"`
	UImage  string `gorm:"column:uimage;not null" gorm:"comment:保存用户头像路径"`
}

type Gz struct { //关注配对数据结构
	Gzid  uint `gorm:"column:gz_id;;not null;primaryKey"`
	Bgzid uint `gorm:"column:bgz_id;;not null;primaryKey"`
}

type EmailMessage struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ScheduleItem struct { //行程返回数据结构
	Date    string `json:"date"`    // 日期，字符串形式
	Content string `json:"content"` // 行程具体内容，字符串形式
}

// UserLevel 用户等级
type UserLevel struct {
	UID   uint `gorm:"column:uid;not null;primaryKey" json:"userID"`
	Level int  `gorm:"column:level;not null" json:"level"`
}

// UserReview 用户评分
type UserReview struct {
	ReviewID    uint `gorm:"column:review_id;primaryKey;autoIncrement" json:"reviewID"`
	ReviewerUID uint `gorm:"column:reviewer_uid;not null" json:"reviewerUID"`
	ReviewedUID uint `gorm:"column:reviewed_uid;not null" json:"reviewedUID"`
	Rating      int  `gorm:"column:rating;not null" json:"rating"`
}
