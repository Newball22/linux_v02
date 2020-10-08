package utils

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbName     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	Lifetime   time.Duration
	IdleConn   int
	OpenConn   int

	AccessKey   string
	SecretKey   string
	Bucket      string
	Qiniuserver string
)

func init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败 err:", err)
		return
	}

	LoadServer(file)
	LoadDatabase(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("123new456")

}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbName = file.Section("database").Key("DbName").MustString("school")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("123456")
	IdleConn = file.Section("database").Key("IdleConn").MustInt(10)
	OpenConn = file.Section("database").Key("OpenConn").MustInt(100)
	Lifetime = file.Section("database").Key("Lifetime").MustDuration(10 * time.Second)

}
func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniuyun").Key("AccessKey").String()
	SecretKey = file.Section("qiniuyun").Key("SecretKey").String()
	Bucket = file.Section("qiniuyun").Key("Bucket").String()
	Qiniuserver = file.Section("qiniuyun").Key("Qiniuserver").String()
}
