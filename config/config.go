package config

import (
	"gopkg.in/ini.v1"
)

//App 参数
type App struct {
	Mode string
	Port string
}

//Mysql 参数
type Mysql struct {
	User     string
	Host     string
	Port     string
	Password string
	Database string
}

//Uxiaowechat  企业微信相关参数
type Uxiaowechat struct {
	CorpID         string //企业微信id
	Corpsecret     string //应用的密钥
	TokenURL       string //获取token
	AgentID        string // 应用id
	CreateChatURL  string //创建群聊
	SendInfoURL    string //发送企业微信群聊
	SendMemberURL  string //推送个人信息
	QYChatInfo     string //获取企业微信登录授权信息
	QYChatUserInfo string //获取企业微信成员信息
}

//IniInfo 获取配置
type IniInfo struct {
	App         App
	Mysql       Mysql
	Uxiaowechat Uxiaowechat
}

//IniParser struct
type IniParser struct {
	confReader *ini.File
}

//IniParseError struct
type IniParseError struct {
	errorInfo string
}

func (e *IniParseError) Error() string { return e.errorInfo }

//Load read ini
func (s *IniParser) Load(configFileName string) error {
	conf, err := ini.Load(configFileName)
	if err != nil {
		s.confReader = nil
		return err
	}
	s.confReader = conf
	return nil
}

//GetString ini string
func (s *IniParser) GetString(section string, key string) string {
	if s.confReader == nil {
		return ""
	}
	r := s.confReader.Section(section)
	if r == nil {
		return ""
	}
	return r.Key(key).String()
}

//SetIniInfo 读取配置缓存
//func (conf *IniInfo) getIniInfo() *IniInfo {
//mode :=
//conf.App = app
//conf.Mysql = mysql
//conf.Uxiaowechat = wechat
//return conf
//}
