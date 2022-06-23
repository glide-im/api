package db

import (
	"fmt"
	"github.com/glide-im/api/internal/config"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func Init() {
	conf := config.MySql
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Db, conf.Charset)
	var err error

	newLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		gormLogger.Config{
			SlowThreshold:             time.Nanosecond, // 慢 SQL 阈值
			LogLevel:                  gormLogger.Warn, // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,           // 禁用彩色打印
		},
	)
	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "im_",
			SingularTable: true,
			//NameReplacer:  nil,
			//NoLowerCase:   false,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	db, err := DB.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetConnMaxIdleTime(time.Minute * 6)

	//DB.LogMode(true)
	//DB.SingularTable(true)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "im_" + defaultTableName
	//}
	initRedis()
}

func initRedis() {

	conf := config.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password:     conf.Password,
		DB:           conf.Db,
		PoolSize:     runtime.NumCPU() * 30,
		MinIdleConns: 10,
	})
	result, err := Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
	logger.D("redis ping: %s", result)
}
