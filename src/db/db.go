// db
package db

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"fmt"
	"strings"
)

var db *sql.DB
var err1 error

func init() {
	db, err1 = sql.Open("mysql", "root:zm12393012998@tcp(10.1.1.56:3306)/golanger?charset=utf8")
	checkErr(err1)
}

func InsertUser(username, userpwd string) bool {
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,userpwd=?")
	defer stmt.Close()
	checkErr(err)

	res, err := stmt.Exec(username, userpwd)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	if id > 0 {
		return true
	}
	return false
}

func QueryUser(username string) (id int, userpwd string) {
	stmt, err := db.Prepare("SELECT id, userpwd FROM userinfo WHERE username=?")
	defer stmt.Close()
	checkErr(err)

	row := stmt.QueryRow(username)
	row.Scan(&id, &userpwd)
	return
}

func QueryAllBlogByUid(uid int) map[int]string {
	stmt, err := db.Prepare("SELECT id,title,content FROM blog_tb WHERE uid=?")
	defer stmt.Close()
	checkErr(err)

	rows, err := stmt.Query(uid)
	defer rows.Close()
	checkErr(err)
	blogmap := make(map[int]string, 10)
	for rows.Next() {
		var id int
		var title, content string
		rows.Scan(&id, &title, &content)
		blogmap[id] = title + "@" + content
	}
	return blogmap
}

func QueryAllBlog() map[int]string {
	stmt, err := db.Prepare("SELECT id,title,content FROM blog_tb ")
	defer stmt.Close()
	checkErr(err)

	rows, err := stmt.Query()
	defer rows.Close()
	checkErr(err)
	blogmap := make(map[int]string, 10)
	for rows.Next() {
		var id int
		var title, content string
		rows.Scan(&id, &title, &content)
		blogmap[id] = title + "@" + content
	}
	return blogmap
}

func QueryBlog(id int) (title, content string) {
	stmt, err := db.Prepare("SELECT title,content FROM blog_tb WHERE id=?")
	defer stmt.Close()
	checkErr(err)

	row := stmt.QueryRow(id)

	row.Scan(&title, &content)
	return
}

func InsertBlog(uid int, title, content string) bool {

	//插入数据
	stmt, err := db.Prepare("INSERT blog_tb SET uid=?, title=?,content=?")
	defer stmt.Close()
	checkErr(err)
	res, err := stmt.Exec(uid, title, strings.TrimSpace(content))
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	if id > 0 {
		return true
	}
	return false
}

func DeleteBlog(bid int) bool {
	stmt, err := db.Prepare("DELETE FROM blog_tb WHERE id=?")
	defer stmt.Close()
	checkErr(err)
	res, err := stmt.Exec(bid)
	checkErr(err)
	fmt.Println(res)
	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
