// @author cold bin
// @date 2023/2/4

package model

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	UserId uint `gorm:"not null"`
	WorkId uint `gorm:"not null"`
}
