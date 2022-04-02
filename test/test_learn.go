package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// 数组和slice
func testArraySlice() {
	a := [3]int8{1, 2, 3}
	a[1] = 4

	fmt.Println(a)

	b := a[1:]
	b[0] = 6

	/*
		在此之前a、b底层数组一样。
		在此之后 append 扩充了slice，导致复出底层数组开辟了新的空间，a和b已经不一样了
	*/

	b = append(b, 10)
	b[1] = 11
	fmt.Println("b: ", b)

	fmt.Println("a: ", a)

	t := fmt.Sprintf("%v", a)

	fmt.Printf("xw: %v", t)

}

func testMap() {

	a := make(map[string]int8)

	fmt.Println(a)

	a["xw"] = 12

	fmt.Println(a["xw"])

}

func bigSlowOperation() {

	defer trace("bigSlowOperation")() // don't forget the extra parentheses

	// ...lots of work…

	time.Sleep(10 * time.Second) // simulate slow operation by sleeping

}

func trace(msg string) func() {

	start := time.Now()

	log.Printf("enter %s", msg)

	return func() {

		log.Printf("exit %s (%s)", msg, time.Since(start))

	}

}

// defer recover忽略异常，程序不终止；如果未提前声明返回值的话，在defer中修改返回值是无效的操作
func passError(s int) (x int, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
			x = 12
		}
	}()
	x = 2 / s
	return x, err
}

type Point struct{ X, Y float64 }
type ColorPoint struct {
	*Point
	color int
}

func (p *Point) ScaleBy(factor float64) {

	p.X *= factor

	p.Y *= factor

}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	var t string
	for i := 0; i < len(s); i++ {
		//fmt.Println(string(s[i]))
		t = string(s[i]) + t
	}

	if s == t {
		return true
	} else {
		return false
	}

}

func spiderUrl(ch chan int) {
	// 根据URL获取资源
	url := "https://www.jb51.net/article/138126.htm"
	res, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	// 读取资源数据 body: []byte
	_, err = ioutil.ReadAll(res.Body)

	// 关闭资源流
	res.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	ch <- 1
}

func main() {

	//ch := make(chan int, 5)

	//for i:=0;i<4;i++{
	//	go spiderUrl(ch)
	//}
	//
	//
	//for i:=0;i<5;i++{
	//	select {
	//	case r := <- ch:
	//			fmt.Println(r)
	//	default:
	//		fmt.Println("zusai")
	//	}
	//
	//}

	//ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		go func(i int) {
			wg.Add(1)
			fmt.Println(i)
			fmt.Println(&i)
			wg.Done()

		}(i)
	}

	wg.Wait()

}

func removeSpace(t *string) *string {
	s := *t
	l := len(s)
	for _, v := range s {
		vNew := string(v)
		if vNew != " " {
			*t = *t + string(v)
		}
	}
	*t = (*t)[l:]
	return t
}
