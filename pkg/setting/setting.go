package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File

	RunMode string

	Host     string
	HTTPPort int

	JWTSecret string
	AppID     string
	AppSecret string
	BusKey    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RunMode").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	Host = sec.Key("HOST").MustString("127.0.0.1")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JWTSecret = sec.Key("JWT_SECRET").MustString("eNQ3a#QyyTXb")
	BusKey = sec.Key("BUSKEY").MustString("eNQ3a#QyyTXb")
	AppID = sec.Key("APPID").String()
	AppSecret = sec.Key("APPSECRET").String()

}
