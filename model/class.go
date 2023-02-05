// @author cold bin
// @date 2023/2/4

package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name      string `gorm:"varchar(25);unique;not null"`
	Introduce string `gorm:"varchar(1000)"`
}
