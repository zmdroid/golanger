// db
package db

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"fmt"
)

var db *sql.DB
var err1 error

func init() {
	db, err1 = sql.Open("mysql", "root:zm12393012998@tcp(10.1.1.56:3306)/golanger?charset=utf8")
	checkErr(err1)
}

func InsertUser(username, userpwd string) {
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,userpwd=?")
	defer stmt.Close()
	checkErr(err)

	res, err := stmt.Exec(username, userpwd)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func QueryUser(username string) (userpwd string) {
	stmt, err := db.Prepare("SELECT userpwd FROM userinfo WHERE username=?")
	defer stmt.Close()
	checkErr(err)

	row := stmt.QueryRow(username)
	row.Scan(&userpwd)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
