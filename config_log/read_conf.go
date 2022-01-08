package config_log

import (
	"flag"
	"github.com/larspensjo/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func getCurrentPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return currentPath
}

func findFIle(currentPath string, fileName string) (string, bool) {
	var filename string
	files, err := ioutil.ReadDir(currentPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename = file.Name()
		if fileName == filename {
			return currentPath + "/" + filename, true
		}
	}
	if "/" == currentPath {
		return currentPath, false
	}
	currentPath = filepath.Dir(currentPath)
	return findFIle(currentPath, fileName)
}

func findConfigFile() (string, bool) {
	currentPath := getCurrentPath()
	filePath, state := findFIle(currentPath, "settings.local.ini")
	if state == false {
		filePath, state = findFIle(currentPath, "settings.ini")
	}
	return filePath, state
}

var (
	//https://studygolang.com/articles/686
	//支持命令行输入格式为-configfile=name, 默认为config.ini
	//配置文件一般获取到都是类型
	filePath, state = findConfigFile()
	configFile      = flag.String("configfile", filePath, "General configuration file")
)

func GetConfigValue(label string, key string) string {
	var value string
	if state == false {
		log.Fatalf("Fail to find %v", filePath)
	}
	cfg, err := config.ReadDefault(*configFile) //读取配置文件，并返回其Config
	if err != nil {
		log.Fatalf("Fail to find %v,%v", *configFile, err)
	}
	if cfg.HasSection(label) { //判断配置文件中是否有section（一级标签）
		value, err := cfg.String(label, key) //根据一级标签section和option获取对应的值
		if err != nil {
			log.Fatalf("Fail to find key: %v", key)
		} else {
			return value
		}
	} else {
		log.Fatalf("Fail to find label: %v", label)
	}
	return value
}
