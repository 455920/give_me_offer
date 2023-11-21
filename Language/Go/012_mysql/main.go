package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 1.连接数据库
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		// 处理连接错误
	}
	defer db.Close()

	// 2.查询数据
	rows, err := db.Query("SELECT * FROM table_name")
	if err != nil {
		// 处理查询错误
	}
	defer rows.Close()

	for rows.Next() {
		// 处理每一行的数据
		var column1 string
		var column2 int
		err := rows.Scan(&column1, &column2)
		if err != nil {
			// 处理扫描错误
		}
		// 处理数据
	}

	// 3.执行插入、更新或删除
	var value1 int
	var value2 int
	result, err := db.Exec("INSERT INTO table_name (column1, column2) VALUES (?, ?)", value1, value2)
	if err != nil {
		// 处理执行错误
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		// 处理获取受影响行数错误
	}
	fmt.Println(affectedRows)

	//
}
