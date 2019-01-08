// Example code for Chapter 5.3 from "Build Web Application with Golang"
// Purpose: Shows how to run simple CRUD operations using a sqlite driver
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const DB_PATH = "./foo.db"

func main() {
	db, err := sql.Open("sqlite3", DB_PATH)
	checkErr(err)
	defer db.Close()

	fmt.Println("Inserting")
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "software developement", time.Now().Format("2006-01-02"))
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
		fmt.Printf("%3v | %6v | %8v | %6v\n", uid, username, department, created)
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
