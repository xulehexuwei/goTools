package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func spiderUrl(ch chan int) {
	// 根据URL获取资源
	url := "https://www.sensedeal.vip/login"
	res, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		//os.Exit(1)
	} else {
		// 读取资源数据 body: []byte
		_, err := ioutil.ReadAll(res.Body)

		// 关闭资源流
		res.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			//os.Exit(1)
		}
	}

	ch <- 1
}

func main() {
	//loopNums := 100000
	//ch := make(chan int, loopNums)
	//
	//for i := 0; i<loopNums; i++{
	//	go spiderUrl(ch)
	//}
	//
	//for i := 0; i<loopNums; i++{
	//	fmt.Println(<- ch)
	//}

	s := "xwe"

	fmt.Printf("%T", s[0])
}
