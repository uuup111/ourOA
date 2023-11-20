package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"ruoyi-go/app/admin/model/constants"
	"ruoyi-go/config"
	"time"
)

// get 方法写入log 文档里面
// post put del 方法 写入数据库
// 错误日志写入数据库

func Logger() gin.HandlerFunc {
	filePath := "log/log"
	//linkName := "latest_log.log"

	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		//panic(R.ReturnFailMsg("err:" + err.Error()))
	}
	logger := logrus.New()

	logger.Out = scr

	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		//retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: constants.TimeFormat,
	})

	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime).Milliseconds()
		spendTime := fmt.Sprintf("%d ms", stopTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
		//	写入数据库

	}
}

// poster logo
func Poster() {
	fg := color.New(color.FgBlue)
	logo := `
  _____                            _             _____         
 |  __ \                          (_)           / ____|        
 | |__) |  _   _    ___    _   _   _   ______  | |  __    ___  
 |  _  /  | | | |  / _ \  | | | | | | |______| | | |_ |  / _ \ 
 | | \ \  | |_| | | (_) | | |_| | | |          | |__| | | (_) |
 |_|  \_\  \__,_|  \___/   \__, | |_|           \_____|  \___/ 
                            __/ |                              
                           |___/                               
` +
		"Author:		OptimisticDevelopers\r\n" +
		"Version:	" + config.ProjectVersion + "\r\n" +
		"MiniGO_SDK: 	" + config.MinGoVersion + "\r\n" +
		"Link: https://gitee.com/OptimisticDevelopers/Ruoyi-Go"
	fg.Println(logo)
}
