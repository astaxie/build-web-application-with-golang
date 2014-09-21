// Example code for Chapter 5.6 from "Build Web Application with Golang"
// Purpose: Shows you have to perform basic CRUD operations for a mongodb driver.
package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name  string
	Phone string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	DB_NAME       = "test"
	DB_COLLECTION = "people"
)

func main() {
	session, err := mgo.Dial("localhost")
	checkError(err)
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(DB_NAME).C(DB_COLLECTION)
	err = c.DropCollection()
	checkError(err)

	ale := Person{"Ale", "555-5555"}
	cla := Person{"Cla", "555-1234"}

	fmt.Println("Inserting")
	err = c.Insert(&ale, &cla)
	checkError(err)

	fmt.Println("Updating")
	ale.Phone = "555-0101"
	err = c.Update(bson.M{"name": "Ale"}, &ale)

	fmt.Println("Querying")
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	checkError(err)
	fmt.Println("Phone:", result.Phone)

	fmt.Println("Deleting")
	err = c.Remove(bson.M{"name": "Ale"})
	checkError(err)
}
