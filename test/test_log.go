package main

import (
	"goTools/config_log"
)

func main() {
	config_log.Logger.Init("test_log")
	defer config_log.Logger.File.Close()
	config_log.Logger.Error.Println("This is a test")
}
