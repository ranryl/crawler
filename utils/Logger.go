package utils

import (
	"crawler/conf"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Loggers is web RequestLogger
var Loggers *logrus.Logger

// GinLogger ...
func GinLogger(logFilePath string, logConfig conf.Log) gin.HandlerFunc {
	fileName := path.Join(logFilePath, logConfig.LogPrefix)
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		PanicError(err)
	}
	Loggers = logrus.New()
	Loggers.SetOutput(src)
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
	logWriter, err := rotatelogs.New(
		fileName+logConfig.TimeFileName,
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(logConfig.SaveMaxAge)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(logConfig.RotationTime)*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Loggers.AddHook(lfHook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		latencyTime := endTime.Sub(startTime)
		Loggers.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}
