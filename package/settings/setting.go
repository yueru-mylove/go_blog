package settings

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init()  {
	var e error
	Cfg, e = ini.Load("conf/app.ini")

	if e != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", e)
	}

	LoadBase()
	LoadServer()
	LoadApp()

}


func LoadBase()  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}


func LoadServer()  {

	server, e := Cfg.GetSection("server")
	if e != nil {
		log.Fatal(2, "Fail to get section 'server': %v", e)
	}

	// log level
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	// port
	HTTPPort = server.Key("HTTP_PORT").MustInt(8000)
	// read timeout
	ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	// write timeout
	WriteTimeout = time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp()  {

	app, e := Cfg.GetSection("app")
	if e != nil {
		log.Fatal(2, "Fail to get section 'app': %v", e)
	}

	JwtSecret = app.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = app.Key("PAGE_SIZE").MustInt(10)
}