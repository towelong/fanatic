package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

func Init() {
	DB = InitMysql()
	// 禁用默认表名的复数形式
	DB.SingularTable(true)
}

func InitMysql() *gorm.DB {
	// "%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.name"),
	)

	openDB, err := gorm.Open("mysql", addr)
	if openDB != nil{
		err = openDB.DB().Ping()
		if err != nil {
			logrus.Info("数据库连接错误~")
		}
		logrus.Info("数据库连接成功~")
	}

	return openDB
}

func Sync() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Permission{})
	DB.AutoMigrate(&UserPermission{})
}

func Close() {
	DB.Close()
}

func OpenLog(){
	DB.LogMode(viper.GetBool("db.log"))  //是否开启log
}