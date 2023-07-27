package config

import (
	"github.com/spf13/viper"
)

var (
	MySql       *MySqlConf
	Redis       *RedisConf
	ApiHttp     *ApiHttpConf
	IMRpcServer *IMRpcServerConf
	Qiniu       *QiniuConfigConf
	Kafka       *KafkaConf
)

type QiniuConfigConf struct {
	QINIU_SK         string
	QINIU_AK         string
	QINIU_BUKET_PATH string
	QINIU_HOST       string
	QINIU_UPLOAD_URL string
	QINIU_UPLOAD_DIR string
}

type KafkaConf struct {
	EnableMessageStorage bool
	Address              []string
}

type ApiHttpConf struct {
	Addr           string
	Port           int
	JwtSecret      string
	IMServerSecret string
}

type IMRpcServerConf struct {
	Addr    string
	Port    int
	Network string
	Etcd    []string
	Name    string
}

type MySqlConf struct {
	Host     string
	Port     int
	Username string
	Password string
	Db       string
	Charset  string
}

type RedisConf struct {
	Host     string
	Port     int
	Password string
	Db       int
}

func MustLoad() {
	err := Load()
	if err != nil {
		panic(err)
	}
}

func Load() error {

	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./_local_config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/.config/")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	c := struct {
		MySql        *MySqlConf
		Redis        *RedisConf
		ApiHttp      *ApiHttpConf
		IMRpcService *IMRpcServerConf
		Qiniu        *QiniuConfigConf
		Kafka        *KafkaConf
	}{}

	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}
	Redis = c.Redis
	MySql = c.MySql
	ApiHttp = c.ApiHttp
	IMRpcServer = c.IMRpcService
	Qiniu = c.Qiniu
	Kafka = c.Kafka

	return err
}
