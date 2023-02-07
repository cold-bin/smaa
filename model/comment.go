// @author cold bin
// @date 2023/2/4

package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	FromWorkId uint `gorm:"not null"`

	FromUserId uint  `gorm:"not null"`
	FromFUser  FUser `gorm:"foreignKey:FromUserId"`

	ToCommentId uint `gorm:"not null;default:0"` // 如果是0，则表示顶级评论

	Content     string `gorm:"not null"`
	IpTerritory string `gorm:"not null"`
}
