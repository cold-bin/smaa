// @author cold bin
// @date 2023/2/4

package model

import (
	"gorm.io/gorm"
	"time"
)

// FUser 前台账户信息表
type FUser struct {
	gorm.Model
	Phone    string `gorm:"type:char(11);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Salt     string `gorm:"type:varchar(255);not null"` // 随机盐，用以增强加密

	SerialNumber string `gorm:"unique;not null;type:varchar(255)"`
	NickName     string `gorm:"type:varchar(255)"`
	Introduce    string `gorm:"type:varchar(3000)"`
	Avatar       string `gorm:"type:varchar(255);not null;default:'http://rplgarpva.bkt.clouddn.com/image/defaut_avatar.jpg'"`
	Background   string `gorm:"type:varchar(255);not null;default:'http://rplgarpva.bkt.clouddn.com/image/default_background.jpg'"`
	Gender       string `gorm:"type:char(1);"` // 1=>男；2=>女
	Residential  string `gorm:"type:varchar(255)"`
	Job          string `gorm:"type:varchar(255)"`
	School       string `gorm:"type:varchar(255)"`

	Height   int    // 身高
	Weight   int    // 体重
	SkinType string `gorm:"type:varchar(255)"` // 肤质
	Birthday time.Time
}
