package orm

import (
	"fmt"
	"go-chatgpt/app/frontend/v1/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Orm *gorm.DB

func Init() {
	mysqlDns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Config.MySQL.User, config.Config.MySQL.Password, config.Config.MySQL.Host, config.Config.MySQL.Port, config.Config.MySQL.DbName, config.Config.MySQL.Charset)
	open, err := gorm.Open(mysql.Open(mysqlDns))
	if err != nil {
		log.Fatal("Database link failure : ", err.Error())
		return
	}

	//err = open.AutoMigrate(models.Users{}, models.UserVipLogs{})
	//if err != nil {
	//	log.Fatal("Database link failure : ", err.Error())
	//	return
	//}
	Orm = open
}
