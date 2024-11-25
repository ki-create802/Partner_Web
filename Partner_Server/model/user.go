package model

type User struct {
	UID     uint   `gorm:"column:uid;not null;primaryKey"`
	UName   string `gorm:"column:uname;type:varchar(45);not null"`
	UKey    string `gorm:"column:ukey;type:varchar(45);not null"`
	UEmail  string `gorm:"column:uemail;type:varchar(45);not null"`
	URemark string `gorm:"column:uremark;type:text"`
	UImage  int    `gorm:"column:uimage;not null" gorm:"comment:保存用户头像id"`
}

type Gz struct {
	Gzid  uint `gorm:"column:gz_id;;not null;primaryKey"`
	Bgzid uint `gorm:"column:bgz_id;;not null;primaryKey"`
}
