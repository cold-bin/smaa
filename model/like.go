// @author cold bin
// @date 2023/2/4

package model

type CommentLike struct {
	ID        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"not null"`
	CommentId uint   `gorm:"not null"`
	IsTop     string `gorm:"not null"` // 是否是顶级评论
}

type WorkLike struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint `gorm:"not null"`
	WorkId uint `gorm:"not null"`
}
