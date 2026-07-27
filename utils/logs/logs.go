// Author: Lutong.li
package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// 初始化
	loggerInit()
}

//type Logs struct {
//}
//
//func (l *Logs) Debug(str string) {}
//
//func (l *Logs) Info(str string) {}
//
//func (l *Logs) Error(str string) {}
//
//func (l *Logs) Warn(str string) {}

// 初始化log日志器的设置
// Debug < Info < Warn < Error < Fatal/Panic
func loggerInit() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//增加行号
	//logrus.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

// 打印debug类型的日志
func Debug(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}

func Info(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Info(msg)
}

func Error(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Error(msg)
}

func Warning(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Warning(msg)
}
