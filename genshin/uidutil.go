package genshin

import (
	"errors"
	"regexp"
)

var uidRe *regexp.Regexp

func init() {
	uidRe, _ = regexp.Compile("[1-257-9][0-9]{8}")
}
func checkUID(uid string) error {
	if uidRe.MatchString(uid) {
		return nil
	}
	return errors.New("非法的UID[" + uid + "]")
}

func server(uid string) (string, string, error) {
	err := checkUID(uid)
	if err != nil {
		return "", "", err
	}
	switch string(uid[0]) {
	case "1":
		fallthrough
	case "2":
		return "cn_gf01", "China", nil
	case "5":
		return "cn_qd01", "China", nil
	case "7":
		return "os_euro", "Europe", nil
	case "8":
		return "os_asia", "Asia", nil
	case "9":
		return "os_cht", "TW/HK/MO", nil
	}
	// never reach here
	return "", "", nil
}
