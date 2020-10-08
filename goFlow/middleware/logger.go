package middleware

//第三方日志logrus的中间件,
import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotaLog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

type Log struct {
	//请求到返回的时间段
	SpendTime string `json:"spend_time"`
	//请求的客户端名
	HostName string `json:"host_name"`
	//请求状态
	StatusCode int `json:"status_code"`
	//客户端IP
	ClientIp string `json:"client_ip"`
	//用户代理
	UserAgent string `json:"user_agent"`
	//数据大小
	DataSize int `json:"data_size"`
	//请求方法
	Method string `json:"method"`
	//请求的URI
	Path string `json:"path"`
}

var (
	err error
	log Log
)

//日志中间件「有日志保存时间，按多长时间分割日志」
func Logger() gin.HandlerFunc {
	filePath := "logData/log"
	linkName := "latest_log.log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logger := logrus.New()
	logger.Out = src //定义logger输出为src

	logger.SetLevel(logrus.DebugLevel) //设置日志等级
	logWriter, _ := rotaLog.New(
		filePath+"%Y%m%d.log",                  //设置生成日志的后缀名
		rotaLog.WithMaxAge(7*24*time.Hour),     //设置日志生命周期：一周
		rotaLog.WithRotationTime(24*time.Hour), //设置分割周期为:24小时分割一次
		rotaLog.WithLinkName(linkName),         //设置一个软链接:直接指向最后一个日志，方便直接看最新的日志j8
	)
	//output interface{}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.PanicLevel: logWriter,
		logrus.ErrorLevel: logWriter,
	}

	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(hook) //文件按时间来分割

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() //洋葱模型的中间件
		stopTime := time.Since(startTime)
		//math.Ceil是向上取整
		log.SpendTime = fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0))) //转换成毫秒为单位的整数
		log.HostName, err = os.Hostname()
		if err != nil {
			log.HostName = "unKnown"
		}
		log.StatusCode = c.Writer.Status()    //状态
		log.ClientIp = c.ClientIP()           //客户端ip
		log.UserAgent = c.Request.UserAgent() //客户端类型{手机还是浏览器}
		log.DataSize = c.Writer.Size()
		if log.DataSize < 0 {
			log.DataSize = 0
		}
		log.Method = c.Request.Method
		log.Path = c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  log.HostName,
			"Status":    log.StatusCode,
			"SpendTime": log.SpendTime,
			"Ip":        log.ClientIp,
			"Method":    log.Method,
			"Path":      log.Path,
			"DataSize":  log.DataSize,
			"UserAgent": log.UserAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String()) //gin框架的错误
		}
		if log.StatusCode >= 500 {
			entry.Error()
		} else if log.StatusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}

	}
}
