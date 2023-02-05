// @author cold bin
// @date 2023/2/4

package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Question string `gorm:"not null"`
	Items    string `gorm:"not null"`
	Answer   string `gorm:"not null"`
}
