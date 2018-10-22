package validator

import "github.com/go-ini/ini"

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

func GetParam(cfg *ini.File, section string, key string) string {
	s, _ := cfg.GetSection(section)
	k, _ := s.GetKey(key)
	return k.String()
}