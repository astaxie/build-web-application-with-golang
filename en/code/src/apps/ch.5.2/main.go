// Example code for Chapter 5.2 from "Build Web Application with Golang"
// Purpose: Use SQL driver to perform simple CRUD operations.
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER     = "user"
	DB_PASSWORD = ""
	DB_NAME     = "test"
)

func main() {
	dbSouce := fmt.Sprintf("%v:%v@/%v?charset=utf8", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("mysql", dbSouce)
	checkErr(err)
	defer db.Close()

	fmt.Println("Inserting")
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "software developement", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id of last inserted row =", id)
	fmt.Println("Updating")
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "row(s) changed")

	fmt.Println("Querying")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username, department, created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | department | created")
		fmt.Printf("%3v | %6v | %6v | %6v\n", uid, username, department, created)
	}

	fmt.Println("Deleting")
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "row(s) changed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
