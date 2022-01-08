package main

import (
	"fmt"
	"goTools/config_log"
)

func testConf() {
	r := config_log.GetConfigValue("mysql", "host")
	fmt.Println(r)
}

func main() {
	testConf()
}
