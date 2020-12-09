package models

import (
	"database/sql"
	"fmt"
	"gin-example/pkg/logging"
	"gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

/* 初始化数据库链接 */
var db *gorm.DB
var sqlDB *sql.DB

type Model struct {
	ID         int            `gorm:"primary_key" json:"id"`
	CreatedOn  int64          `json:"created_on"`
	ModifiedOn int64          `json:"modified_on"`
	DeletedOn  gorm.DeletedAt `json:"deleted_on"`
}

func Setup() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)

	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	)

	// 日志相关
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,          // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 所有 table 前缀
			SingularTable: true,        // 最后 table 不加s
		},
		Logger: newLogger,
	})

	if err != nil {
		logging.Error("models.Setup err: %v", err)
	}

	//gorm.DefaultTableNameHandler = func() {}
	sqlDB, err = db.DB()
	if err != nil {
		logging.Error("get db.BD() error: %v", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 创建表
	//_ = db.AutoMigrate(&Article{}, &Auth{}, &Tag{})
}

func Close() {
	sqlDB.Close()
}
