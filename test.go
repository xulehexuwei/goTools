package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Engin
	router := gin.Default()
	router.GET("/hello", hello) // hello函数处理"/hello"请求
	// 指定地址和端口号
	router.Run(":9090")
}

func hello(context *gin.Context) {
	println(">>>> hello function start <<<<")

	var b []map[string]interface{}

	b = make([]map[string]interface{}, 0)

	b = append(b, map[string]interface{}{"xw": "xw", "count": 23})
	fmt.Println(b)

	context.JSON(http.StatusOK, gin.H{
		"code":    &b,
		"success": true,
	})
}
