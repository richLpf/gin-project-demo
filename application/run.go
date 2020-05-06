package application

import (
	"flag"
	"fmt"
	"myapp/config"
	"myapp/dbs"
	apiRouter "myapp/router"
	"os"
)

//Run 启动app
func Run() {
	mode := flag.String("mode", "dev", "eventment")
	flag.Parse()
	mysql, err := getMysql(*mode)
	if err != nil {
		panic(err)
	}
	dbs.RunMysql(mysql)
	app, err := getApp(*mode)
	fmt.Println("app", app)
	if err != nil {
		panic(err)
	}
	router := apiRouter.Router(app)
	router.Run(app.Port)
}

func readIni(mode string) (iniParser config.IniParser, err error) {
	confFileName := "dev.ini"
	if mode == "prod" {
		confFileName = "prod.ini"
	}
	//iniParser = conf.IniParser{}
	dir, _ := os.Getwd()
	pathName := dir + `/config/` + confFileName
	err = iniParser.Load(pathName)
	return iniParser, err
}

func getApp(mode string) (app config.App, err error) {
	iniParser, err := readIni(mode)
	if err != nil {
		return app, err
	}
	app = config.App{
		Mode: mode,
		Port: iniParser.GetString("app", "port"),
	}
	return app, nil
}

func getMysql(mode string) (mysql config.Mysql, err error) {
	iniParser, err := readIni(mode)
	if err != nil {
		return mysql, err
	}
	mysql = config.Mysql{
		User:     iniParser.GetString("mysql", "user"),
		Password: iniParser.GetString("mysql", "password"),
		Host:     iniParser.GetString("mysql", "host"),
		Port:     iniParser.GetString("mysql", "port"),
		Database: iniParser.GetString("mysql", "database"),
	}
	return mysql, err
}
