// @author cold bin
// @date 2023/2/4

package model

import "gorm.io/gorm"

type Work struct {
	gorm.Model

	UserId    uint  `gorm:"not null"` // 创作人
	FromFUser FUser `gorm:"foreignKey:UserId"`

	Type        string `gorm:"type:char(1);not null"` // 类型 1=>富文本 2=>视频
	ClassId     uint   `gorm:"not null"`              // 类别
	Title       string `gorm:"not null"`              // 标题
	EntryUrl    string `gorm:"not null"`              // 内容链接
	IpTerritory string `gorm:"not null"`              // ip属地
	ReadNum     string `gorm:"not null;default:0"`    // 阅读量
}
