package validator

import (
	"errors"
	"github.com/go-ini/ini"
)

const defaultLang = "zh_CN"

var dict = make(map[string]*ini.File)

// loading language
func LoadLang(lang string, file string) {
	f, err := ini.Load(file)
	if err != nil {
		panic("load " + file + " failed")
	}
	dict[lang] = f
}

func getParam(cfg *ini.File, section string, key string) (string, error) {
	s, err1 := cfg.GetSection(section)
	if err1 != nil {
		return "", errors.New("tpl " + section + " not exist")
	}
	k, err2 := s.GetKey(key)
	if err2 != nil {
		return "", errors.New("dict " + key + " not exist")
	}
	return k.String(), nil
}

func GetLang(lang string) *ini.File {
	return dict[lang]
}
