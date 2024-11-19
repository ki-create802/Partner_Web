package model

type User struct {
	UID     string `gorm:"type:varchar(45);not null;primaryKey"`
	UName   string `gorm:"type:varchar(45);not null"`
	UKey    string `gorm:"type:varchar(45);not null"`
	UEmail  string `gorm:"type:varchar(45);not null"`
	URemark string `gorm:"type:text"`
	UImage  uint   `gorm:"not null" gorm:"comment:保存用户头像id"`
}
