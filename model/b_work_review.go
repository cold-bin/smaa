// @author cold bin
// @date 2023/2/4

package model

import "gorm.io/gorm"

// BWorkReview 作品审核表
type BWorkReview struct {
	gorm.Model
	UserId      uint   `gorm:"not null"`
	ClassId     uint   `gorm:"not null"`
	Type        string `gorm:"varchar(255);not null"`
	Title       string `gorm:"varchar(255);not null"`
	EntryUrl    string `gorm:"varchar(255);not null"`
	IpTerritory string `gorm:"varchar(25);not null"`
}
