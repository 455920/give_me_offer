在Go语言中，可以使用第三方库来操作MySQL数据库。其中，database/sql是Go语言标准库中提供的数据库操作接口，而go-sql-driver/mysql是一个常用的MySQL驱动程序。

以下是一个简单的示例，演示如何使用Go操作MySQL数据库：

安装依赖：首先，需要安装go-sql-driver/mysql库。可以使用以下命令进行安装：
```
go get -u github.com/go-sql-driver/mysql
```
导入依赖：在Go程序中，导入database/sql和github.com/go-sql-driver/mysql包：
```
import (
"database/sql"
_ "github.com/go-sql-driver/mysql"
)
```
连接数据库：使用sql.Open函数来连接MySQL数据库，并返回一个*sql.DB对象：
```
db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
if err != nil {
// 处理连接错误
}
defer db.Close()
```
在上述代码中，需要将username、password和database_name替换为实际的数据库连接信息。

执行查询：使用db.Query函数执行SQL查询，并获取结果集：
```
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
```
在上述代码中，需要将table_name替换为实际的表名，并根据需要定义变量来接收查询结果的列。

执行插入、更新或删除：使用db.Exec函数执行插入、更新或删除操作：
```
result, err := db.Exec("INSERT INTO table_name (column1, column2) VALUES (?, ?)", value1, value2)
if err != nil {
// 处理执行错误
}
affectedRows, err := result.RowsAffected()
if err != nil {
// 处理获取受影响行数错误
}
```
在上述代码中，需要将table_name替换为实际的表名，并根据需要传递参数来替换?占位符。

这只是一个简单的示例，演示了如何使用Go操作MySQL数据库。实际应用中，可能还需要处理错误、事务、连接池等更复杂的情况。