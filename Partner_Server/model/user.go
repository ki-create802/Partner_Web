package model

type User struct {
	UID     uint   `gorm:"column:uid;not null;primaryKey"`
	UName   string `gorm:"column:uname;type:varchar(45);not null"`
	UKey    string `gorm:"column:ukey;type:varchar(45);not null" json:"Password"`
	UEmail  string `gorm:"column:uemail;type:varchar(45);not null" json:"Email"`
	URemark string `gorm:"column:uremark;type:text"`
	UImage  string `gorm:"column:uimage;not null" gorm:"comment:保存用户头像路径"`
}

type Gz struct {
	Gzid  uint `gorm:"column:gz_id;;not null;primaryKey"`
	Bgzid uint `gorm:"column:bgz_id;;not null;primaryKey"`
}

type EmailMessage struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
