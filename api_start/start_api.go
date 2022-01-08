package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goTools/db"
	"net/http"
	"time"
)

func chanData(cha chan []map[string]string) {
	fmt.Println("start")
	time.Sleep(6 * time.Second)
	fmt.Println("end")
	sql := "select * from qa_xiaomu limit 10;"
	cha <- db.GetDataByQuery("mysql", "qa_log", sql)

}

//
//func getMysqlData(context *gin.Context) {
//	cha := make(chan []map[string]string, 10)
//	sql := "select * from qa_xiaomu limit 10;"
//	go time.Sleep(6 * time.Second)
//	cha <- db.GetDataByQuery("mysql", "qa_log", sql)
//	context.JSON(http.StatusOK, gin.H{
//		"data":    &data,
//		"success": true,
//	})
//}

func getMysqlData(context *gin.Context) {
	cha := make(chan []map[string]string, 10)
	go chanData(cha)
	context.JSON(http.StatusOK, gin.H{
		"data":    <-cha,
		"success": true,
	})
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

func main() {
	// 注册路由
	router := gin.Default()
	router.GET("", hello)
	router.GET("/data", getMysqlData)

	//router.GET("/data", func(c *gin.Context) {
	//	cCp := c.Copy()
	//	go func() {
	//		sql := "select * from qa_xiaomu limit 10;"
	//		data := db.GetDataByQuery("mysql", "qa_log", sql)
	//		//time.Sleep(6 * time.Second)
	//		cCp.JSON(http.StatusOK, gin.H{
	//			"data":  &data,
	//			"success": true,
	//		})
	//	}()
	//})

	// 指定地址和端口号
	router.Run(":8000")
}
