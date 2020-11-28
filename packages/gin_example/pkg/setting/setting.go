package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cig          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cig, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("加载初始化文件 'conf/app.ini' 文件失败: %v", err)
	}
}

func LoadBase() {
	RunMode = Cig.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cig.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cig.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
