package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"goFlow/utils"
)

var (
	DB  *gorm.DB
	err error
)

func InitDb() {
	DB, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	if err != nil {
		logrus.Panic(err)
	}
	//数据库
	DB.SingularTable(true)
	//数据自动迁移
	DB.AutoMigrate(&User{}, &Article{}, &Category{})

	/*gorm虚拟出来的连接池*/
	//最大空闲连接数
	DB.DB().SetMaxIdleConns(utils.IdleConn)
	//最大连接数
	DB.DB().SetMaxOpenConns(utils.OpenConn)
	//设置连接最大的可复用时间
	DB.DB().SetConnMaxLifetime(utils.Lifetime)
	//上线的时候需要close
	//defer db.Close()
}
