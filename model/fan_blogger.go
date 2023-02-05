// @author cold bin
// @date 2023/2/4

package model

import "gorm.io/gorm"

type FanBlogger struct {
	gorm.Model
	FanId     uint `gorm:"not null"`
	BloggerId uint `gorm:"not null"`
}
