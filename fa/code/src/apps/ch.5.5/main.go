// Example code for Chapter 5.5
// Purpose is to show to use BeeDB ORM for basic CRUD operations for sqlite3
package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var orm beedb.Model

type Userinfo struct {
	Uid        int `beedb:"PK"`
	Username   string
	Department string
	Created    string
}

const DB_PATH = "./foo.db"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func getTimeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func insertUsingStruct() int64 {
	fmt.Println("insertUsingStruct()")
	var obj Userinfo
	obj.Username = "Test Add User"
	obj.Department = "Test Add Department"
	obj.Created = getTimeStamp()
	checkError(orm.Save(&obj))
	fmt.Printf("%+v\n", obj)
	return int64(obj.Uid)
}
func insertUsingMap() int64 {
	fmt.Println("insertUsingMap()")
	add := make(map[string]interface{})
	add["username"] = "astaxie"
	add["department"] = "cloud develop"
	add["created"] = getTimeStamp()
	id, err := orm.SetTable("userinfo").Insert(add)
	checkError(err)
	fmt.Println("Last row inserted id =", id)
	return id
}

func getOneUserInfo(id int64) Userinfo {
	fmt.Println("getOneUserInfo()")
	var obj Userinfo
	checkError(orm.Where("uid=?", id).Find(&obj))
	return obj
}

func getAllUserInfo(id int64) []Userinfo {
	fmt.Println("getAllUserInfo()")
	var alluser []Userinfo
	checkError(orm.Limit(10).Where("uid>?", id).FindAll(&alluser))
	return alluser
}

func updateUserinfo(id int64) {
	fmt.Println("updateUserinfo()")
	var obj Userinfo
	obj.Uid = int(id)
	obj.Username = "Update Username"
	obj.Department = "Update Department"
	obj.Created = getTimeStamp()
	checkError(orm.Save(&obj))
	fmt.Printf("%+v\n", obj)
}

func updateUsingMap(id int64) {
	fmt.Println("updateUsingMap()")
	t := make(map[string]interface{})
	t["username"] = "updateastaxie"
	//update one
	// id, err := orm.SetTable("userinfo").SetPK("uid").Where(2).Update(t)
	//update batch
	lastId, err := orm.SetTable("userinfo").Where("uid>?", id).Update(t)
	checkError(err)
	fmt.Println("Last row updated id =", lastId)
}

func getMapsFromSelect(id int64) []map[string][]byte {
	fmt.Println("getMapsFromSelect()")
	//Original SQL Backinfo resultsSlice []map[string][]byte
	//default PrimaryKey id
	c, err := orm.SetTable("userinfo").SetPK("uid").Where(id).Select("uid,username").FindMap()
	checkError(err)
	fmt.Printf("%+v\n", c)
	return c
}

func groupby() {
	fmt.Println("groupby()")
	//Original SQL Group By
	b, err := orm.SetTable("userinfo").GroupBy("username").Having("username='updateastaxie'").FindMap()
	checkError(err)
	fmt.Printf("%+v\n", b)
}

func joinTables(id int64) {
	fmt.Println("joinTables()")
	//Original SQL Join Table
	a, err := orm.SetTable("userinfo").Join("LEFT", "userdetail", "userinfo.uid=userdetail.uid").Where("userinfo.uid=?", id).Select("userinfo.uid,userinfo.username,userdetail.profile").FindMap()
	checkError(err)
	fmt.Printf("%+v\n", a)
}

func deleteWithUserinfo(id int64) {
	fmt.Println("deleteWithUserinfo()")
	obj := getOneUserInfo(id)
	id, err := orm.Delete(&obj)
	checkError(err)
	fmt.Println("Last row deleted id =", id)
}

func deleteRows() {
	fmt.Println("deleteRows()")
	//original SQL delete
	id, err := orm.SetTable("userinfo").Where("uid>?", 2).DeleteRow()
	checkError(err)
	fmt.Println("Last row updated id =", id)
}

func deleteAllUserinfo(id int64) {
	fmt.Println("deleteAllUserinfo()")
	//delete all data
	alluser := getAllUserInfo(id)
	id, err := orm.DeleteAll(&alluser)
	checkError(err)
	fmt.Println("Last row updated id =", id)
}
func main() {
	db, err := sql.Open("sqlite3", DB_PATH)
	checkError(err)
	orm = beedb.New(db)
	var lastIdInserted int64

	fmt.Println("Inserting")
	lastIdInserted = insertUsingStruct()
	insertUsingMap()

	a := getOneUserInfo(lastIdInserted)
	fmt.Println(a)

	b := getAllUserInfo(lastIdInserted)
	fmt.Println(b)

	fmt.Println("Updating")
	updateUserinfo(lastIdInserted)
	updateUsingMap(lastIdInserted)

	fmt.Println("Querying")
	getMapsFromSelect(lastIdInserted)
	groupby()
	joinTables(lastIdInserted)

	fmt.Println("Deleting")
	deleteWithUserinfo(lastIdInserted)
	deleteRows()
	deleteAllUserinfo(lastIdInserted)
}
