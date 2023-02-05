// @author cold bin
// @date 2023/2/4

package _mysql

import (
	"eliminate-male-appearance-anxiety/global"
	"eliminate-male-appearance-anxiety/model"
	"fmt"
	"github.com/cold-bin/goutil/general/mlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitMysql() error {
	_mysql := global.Set.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		_mysql.UserName,
		_mysql.Password,
		_mysql.Host,
		_mysql.Port,
		_mysql.Db,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(_mysql.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(_mysql.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(_mysql.ConnMaxLifetime) * time.Hour)

	mlog.Debugf("设置了连接可复用的最大时间: %d hour", time.Duration(_mysql.ConnMaxLifetime)*time.Hour)

	global.GDB = db

	// 迁移表数据
	if err = migrate(db); err != nil {
		return err
	}

	return nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.FUser{},
		&model.BUser{},
		&model.BWorkReview{},
		&model.Collection{},
		&model.CommentSecondary{},
		&model.CommentTop{},
		&model.FanBlogger{},
		&model.Question{},
		&model.Work{},
		&model.Class{},
		&model.CommentLike{},
		&model.WorkLike{},
	)
}
