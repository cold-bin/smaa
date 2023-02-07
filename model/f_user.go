// @author cold bin
// @date 2023/2/4

package model

import (
	"eliminate-male-appearance-anxiety/global"
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

	IsVisibleCollection bool `gorm:"default:0;not null"` // 默认收藏都是不可见
	IsVisibleLike       bool `gorm:"default:0;not null"` // 默认点赞都是不可见

	Height   uint16 // 身高
	Weight   uint16 // 体重
	SkinType string `gorm:"type:varchar(255)"` // 肤质
	Birthday time.Time
}

func (u *FUser) Create() error {
	return global.GDB.Create(u).Error
}

// FindAllByPhone 提供phone
func (u *FUser) FindAllByPhone() error {
	return global.GDB.Where("phone = ?", u.Phone).Find(u).Error
}

// FindFieldByPhone 需要手机号，并且需要指定查找的字段名
func (u *FUser) FindFieldByPhone(fields ...string) error {
	return global.GDB.Select(fields).Where("phone = ?", u.Phone).Find(u).Error
}

// UpdateFieldByPhone 更新字段
func (u *FUser) UpdateFieldByPhone(fields ...string) error {
	return global.GDB.Model(u).Select(fields).Where("phone = ?", u.Phone).Updates(u).Error
}
