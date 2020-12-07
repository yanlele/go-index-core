package setting

import (
	"gin-example/pkg/logging"
	"github.com/go-ini/ini"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func init() {
	appConfig, err := ini.Load("conf/app.ini")
	if err != nil {
		logging.Fatal("加载初始化文件 'conf/app.ini' 文件失败: %v", err)
	}

	err = appConfig.Section("app").MapTo(AppSetting)
	if err != nil {
		logging.Fatal("config mapTo AppSetting err : ", err.Error())
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = appConfig.Section("server").MapTo(ServerSetting)
	if err != nil {
		logging.Fatal("config mapTo ServerSetting err: ", err.Error())
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = appConfig.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		logging.Fatal("config mapTo DatabaseSetting err: ", err.Error())
	}
}

//func LoadBase() {
//	RunMode = Cig.Section("").Key("RUN_MODE").MustString("debug")
//}
//
//func LoadServer() {
//	sec, err := Cig.GetSection("server")
//	if err != nil {
//		logging.Fatal("Fail to get section 'server': %v", err)
//	}
//	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
//	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
//	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
//}
//
//func LoadApp() {
//	sec, err := Cig.GetSection("app")
//	if err != nil {
//		logging.Fatal("Fail to get section 'app': %v", err)
//	}
//	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
//	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
//}
