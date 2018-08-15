package conf

import (
	"log"

	"github.com/go-ini/ini"
	"time"
)

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	DB          int
}

var RedisSetting = make(map[string] *Redis)

var cfg *ini.File

func Setup()  {
	var err error
	RedisSetting["default"] = &Redis{}
	RedisSetting["exam"] = &Redis{}
	RedisSetting["student"] = &Redis{}

	cfg, err = ini.Load("conf/app.ini")
	if err != nil{
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting["default"])
	mapTo("redis-exam", RedisSetting["exam"])
	mapTo("redis-student", RedisSetting["student"])
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil{
		log.Fatalf("Cfg.MapTo Setting err: %v", err)
	}

}