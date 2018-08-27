package Mgo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)
type Person struct {
	Name  string
	Phone string
}

func main() {
	session := Mongo()
	c := session.DB("test").C("people")
	err := c.Insert(&Person{"superWang", "13478808311"},
		&Person{"David", "15040268074"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "superWang"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", result.Name)
	fmt.Println("Phone:", result.Phone)
}

func Mongo() *mgo.Session{
	session, err := mgo.Dial("mongodb://root:ckzc2018abcd@dds-2ze73f189700e204-pub.mongodb.rds.aliyuncs.com:3717/admin")
	if err != nil {
		panic(err)
	}else{
		//defer session.Close()

		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)
		return session
	}

}