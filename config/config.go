package config

import "github.com/spf13/viper"

var MysqlDB string
var RedisDB string
var ElasticsearchDB string
var IdentityConfig IdentityConf
var AliPayConfig AliPayConf

type IdentityConf struct {
	SecretID  string `yaml:"secret_id"`
	SecretKey string `yaml:"secret_key"`
}
type AliPayConf struct {
	AppId      string `yaml:"app_id"`
	PrivateKey string `yaml:"private_key"`
}

func init() {
	path := "D:/go_daima/gao2/month/config/config.yaml"
	viper.SetConfigFile(path)
	viper.ReadInConfig()
	MysqlDB = viper.GetString("Mysql.DB")
	RedisDB = viper.GetString("Redis.DB")
	ElasticsearchDB = viper.GetString("Elasticsearch.DB")
	IdentityConfig = IdentityConf{
		SecretID:  viper.GetString("Identity.SecretID"),
		SecretKey: viper.GetString("Identity.SecretKey"),
	}
	AliPayConfig = AliPayConf{
		AppId:      viper.GetString("AliPay.AppId"),
		PrivateKey: viper.GetString("AliPay.PrivateKey"),
	}
}
