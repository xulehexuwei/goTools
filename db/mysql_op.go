package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //这个包必须导入，才能和mysql连接成功
	"goTools/config_log"
)

func ConMysql(label, dbname string) (*sql.DB, error) {
	host := config_log.GetConfigValue(label, "host")
	port := config_log.GetConfigValue(label, "port")
	user := config_log.GetConfigValue(label, "user")
	pass := config_log.GetConfigValue(label, "pass")
	if dbname == "" {
		dbname = config_log.GetConfigValue(label, "db")
	}
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, dbname)
	//fmt.Println(dataSourceName)
	//db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")
	conDb, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("链接失败")
		return nil, err
	} else {
		fmt.Println("链接数据库成功...........已经打开")
		return conDb, nil
	}
}

func GetDataByQuery(label, dbname, sql string) []map[string]string {
	//打开数据库
	db, error := ConMysql(label, dbname)
	if error != nil {
		return nil
	}
	defer db.Close()
	////获取所有数据
	rows, _ := db.Query(sql)
	defer rows.Close()
	//返回所有列
	cols, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	result := make([]map[string]string, 0)
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result = append(result, row)
	}
	return result
}
