package confini

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
