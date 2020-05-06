package dbs

import (
	"fmt"
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

func init() {
	//本地启动了3306的数据库
	DB, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/web?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)                                                   // 创建表不会默认变成复数
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}) //自动迁移数据库
	fmt.Println("db", DB)
	DB.DB().Ping()
}
