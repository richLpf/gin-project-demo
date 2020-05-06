package dbs

import (
	"fmt"
	"myapp/config"
	"myapp/model"

	"github.com/jinzhu/gorm"

	//mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//DB 数据库
	DB  *gorm.DB
	err error
)

//RunMysql 启动mysql
func RunMysql(mysql config.Mysql) {
	connectSQL := mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.Database + "?charset=utf8mb4,utf8&parseTime=true&loc=Local"
	fmt.Println("connet", connectSQL)
	DB, err = gorm.Open("mysql", connectSQL)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB", DB)
	DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(
		&model.User{},
	)
	DB.DB().Ping()
}
