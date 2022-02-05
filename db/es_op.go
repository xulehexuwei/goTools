package db

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"goTools/config_log"
	"log"
	"os"
	"time"
)

func EsClient(label string) (*elastic.Client, error) {
	url := config_log.GetConfigValue(label, "host")
	user := config_log.GetConfigValue(label, "user", "")
	pass := config_log.GetConfigValue(label, "pass", "")
	// 创建Client, 连接ES
	client, err := elastic.NewClient(
		// Go无法连接docker中es，代码设置sniff 为false
		elastic.SetSniff(false),
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL(url),
		// 基于http base auth验证机制的账号和密码
		elastic.SetBasicAuth(user, pass),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置请求失败最大重试次数
		elastic.SetMaxRetries(5),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))

	return client, err
}

func EsQueryByMatch(label, index, column, text string) *elastic.SearchResult {
	// label是settings.ini中es的连接配置的标签
	// column是es中要查询的字段名称， text是输入的检索内容
	client, error := EsClient(label)
	fmt.Println(error)
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	matchQuery := elastic.NewMatchQuery(column, text)

	searchResult, _ := client.Search().
		Index(index).      // 设置索引名
		Query(matchQuery). // 设置查询条件
		//Sort("Created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(0).      // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(40).     // 设置分页参数 - 每页大小
		Pretty(true). // 查询结果返回可读性较好的JSON格式
		Do(ctx)       // 执行请求

	return searchResult
}
