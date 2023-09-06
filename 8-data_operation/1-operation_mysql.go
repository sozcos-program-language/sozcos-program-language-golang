/*
1-operation_mysql.go: 操作mysql数据库
*/
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
执行以下命令行引入开源MySQL库:
    go get github.com/go-sql-driver/mysql
    go get github.com/jmoiron/sqlx
*/

var Db *sqlx.DB

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func init() {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	database, err := sqlx.Open("mysql", "root:root@tcp(192.168.100.115:3306)/go-database")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func insertPerson(c Person) {

	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", c.Username, c.Sex, c.Email)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}

func selectPerson(id int8) {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
}

/*
		go MySQL事务应用
		Db.Begin()        开始事务
	    Db.Commit()        提交事务
	    Db.Rollback()     回滚事务
*/
func transaction() {
	begin, err2 := Db.Begin()
	if err2 != nil {
		fmt.Println("开启事务异常: >>>>>>>>>>>> err2 =", err2)
		return
	}

	_, err := Db.Exec("update person set email = ? where user_ida = ?", "poozs@163.com", 2)
	if err != nil {
		fmt.Println("MySQL事务回滚: >>>>>>>>>>>> err =", err)
		begin.Rollback()
		return
	}

	begin.Commit()
}

func operationMysql() {
	p := Person{Username: "小明1", Sex: "女", Email: "poozs@qq.com"}
	insertPerson(p)
	selectPerson(3)
	transaction()
}
