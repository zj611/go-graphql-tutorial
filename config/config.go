/**
 * @Date: 2020/11/22 下午2:01
 */
package config

import (
	"github.com/caarlos0/env"
	"github.com/caarlos0/env/parsers"
	"net/url"
)

type Config struct {

	//MysqlURL                    url.URL `env:"MYSQL_URL" envDefault:"mysql://fccvapp:sf123456fccvapp@appmanager-m.dbsit.sfcloud.local:3306/smartimage"`
	MysqlURL url.URL `env:"MYSQL_URL" envDefault:"mysql://root:123@0.0.0.0:3306/data"`
	//MysqlURLMaster              url.URL `env:"MysqlURLMaster" envDefault:"root:root@tcp(0.0.0.0:3308)/test?charset=utf8&parseTime=true"`
	//MysqlURLSlave               url.URL `env:"MysqlURLSlave" envDefault:"root:root@tcp(0.0.0.0:3307)/test?charset=utf8&parseTime=true"`

}

func ReadEnv() *Config {
	cfg := Config{}
	if err := env.ParseWithFuncs(&cfg, env.CustomParsers{
		parsers.URLType: parsers.URLFunc,
	}); err != nil {
		panic(err)
	}
	return &cfg
}
