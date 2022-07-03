package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"model"`
	Username   string `gorm:"type:varchar(100);default:null" json:"username"`
	Password   string `gorm:"type:varchar(100);default:null" json:"password"`
	Fullname   string `gorm:"type:varchar(100);default:null" json:"fullname"`
}
