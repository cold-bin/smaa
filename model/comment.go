// @author cold bin
// @date 2023/2/4

package model

import (
	"gorm.io/gorm"
)

type CommentTop struct {
	gorm.Model
	FromWorkId uint `gorm:"not null"`

	FromUserId uint  `gorm:"not null"`
	FromFUser  FUser `gorm:"foreignKey:FromUserId"`

	IsExistSec  string `gorm:"type:char(1);not null;default:'0'"` // 是否存在子评论,默认不存在（1表示存在,0表示不存在）
	Content     string `gorm:"not null"`
	IpTerritory string `gorm:"not null"`
}

type CommentSecondary struct {
	gorm.Model
	FromWorkId uint `gorm:"not null"`

	FromUserId uint  `gorm:"not null"`
	FromFUser  FUser `gorm:"foreignKey:FromUserId"`

	ToUserId uint  `gorm:"not null"`
	ToFUser  FUser `gorm:"foreignKey:ToUserId"`

	Content     string `gorm:"not null"`
	IpTerritory string `gorm:"not null"`
}
