package mysqlHelper

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/url"
	"strings"
	"time"
)

func NewGormDB(url *url.URL) *gorm.DB {
	cfg := cfgFromURL(url)
	addr := cfg.FormatDSN()

	db, err := gorm.Open("mysql", addr)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}

func cfgFromURL(url *url.URL) *mysql.Config {
	cfg := mysql.Config{
		Addr:                 url.Host,
		DBName:               strings.Replace(url.Path, "/", "", 1),
		Loc:                  time.Local,
		Net:                  "tcp",
		Params:               map[string]string{"parseTime": "true"},
		AllowNativePasswords: true,
	}
	if url.Port() == "" {
		cfg.Addr = cfg.Addr + ":3306"
	}
	if url.User != nil {
		cfg.User = url.User.Username()
		if pwd, ok := url.User.Password(); ok {
			cfg.Passwd = pwd
		}
	}
	return &cfg
}
