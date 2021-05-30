/**
 * created by
 * @project go-improve
 * @author frankstar
 * @date 2021/5/30
 * @contact frankstarye@tencent.com
 **/

package mysql

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	name string
	email string
	phone string
	age int
	sex int
	addtime time.Time
}

func connect() *sql.DB {
	//-------1、打开数据库--------
	db, err := sql.Open("mysql", "root:frankstar@tcp(127.0.0.1:3306)/go_study?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return db
	}
	//延迟到函数结束关闭链接
	//defer db.Close()

	return db
}

func Select(id int) User {
	//-------2、查询单条数据--------
	//定义接收数据的结构
	var u User
	//-------1、打开数据库--------
	db, err := sql.Open("mysql", "root:frankstar@tcp(127.0.0.1:3306)/go_study?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return u
	}
	//执行单条查询
	rows := db.QueryRow("select * from user_tb where id = ?", id)
	rows.Scan(&u.Id, &u.name, &u.age, &u.sex, &u.addtime)
	fmt.Println("单条数据结果：", u)

	defer db.Close()
	return u
}


