package config_log

import (
	"io"
	"log"
	"os"
	"path"
)

type Logg struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	File    *os.File
}

// 初始化日志文件位置
func (logger *Logg) Init(filename string) {
	logPath := GetConfigValue("settings", "log_path")
	logFile := path.Join(logPath, filename+".log")
	File, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Log file does not exist：", err)
	}
	logger.Info = log.New(io.MultiWriter(os.Stderr, File), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warning = log.New(io.MultiWriter(os.Stderr, File), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(io.MultiWriter(os.Stderr, File), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

var Logger Logg
