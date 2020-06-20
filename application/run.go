package application

import (
	"flag"
	"fmt"
	"myapp/config"
	"myapp/dbs"
	apiRouter "myapp/router"
	_ "myapp/services/demo"
	"myapp/services/wechat"
	"os"
	"reflect"

	"github.com/richLpf/goutils/utils"
	uuid "github.com/satori/go.uuid"
)

//Run 启动app
func Run() {
	utils.PrintStr()
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
	Uxiaowechat, err := getQyChartInfo(*mode)
	if err != nil {
		panic(err)
	}
	wechat.GetQyInfo(Uxiaowechat)
	fmt.Println("Uxiaowechat", Uxiaowechat)
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	fmt.Println("uuid type", reflect.TypeOf(u1))

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

func getQyChartInfo(mode string) (wechat config.Uxiaowechat, err error) {
	iniParser, err := readIni(mode)
	if err != nil {
		return wechat, err
	}
	wechat = config.Uxiaowechat{
		CorpID:         iniParser.GetString("Uxiaowechat", "CorpID"),
		Corpsecret:     iniParser.GetString("Uxiaowechat", "Corpsecret"),
		TokenURL:       iniParser.GetString("Uxiaowechat", "TokenURL"),
		AgentID:        iniParser.GetString("Uxiaowechat", "AgentID"),
		CreateChatURL:  iniParser.GetString("Uxiaowechat", "CreateChatURL"),
		SendInfoURL:    iniParser.GetString("Uxiaowechat", "SendInfoURL"),
		SendMemberURL:  iniParser.GetString("Uxiaowechat", "SendInfoURL"),
		QYChatInfo:     iniParser.GetString("Uxiaowechat", "QYChatInfo"),
		QYChatUserInfo: iniParser.GetString("Uxiaowechat", "QYChatUserInfo"),
	}
	return wechat, err
}
