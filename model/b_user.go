// @author cold bin
// @date 2023/2/4

package model

import (
	"eliminate-male-appearance-anxiety/global"
	"gorm.io/gorm"
)

type BUser struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Account  string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Salt     string `gorm:"type:varchar(255);not null"`
	Level    int    `gorm:"default:0"`
}

// Create 插入新纪录
func (u *BUser) Create() error {
	return global.GDB.Create(u).Error
}
