package setting

import (
	"github.com/go-ini/ini"
	"log"
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

func Setup() {
	appConfig, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("加载初始化文件 'conf/app.ini' 文件失败: %v", err)
	}

	err = appConfig.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("config mapTo AppSetting err : %v", err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = appConfig.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("config mapTo ServerSetting err: %v", err)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = appConfig.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("config mapTo DatabaseSetting err: %v", err)
	}
}
