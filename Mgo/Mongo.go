package Mgo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
	"time"
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
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"192.168.31.232"},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  "test",
		Source:    "admin",
		Username:  "",
		Password:  "",
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		panic(err)
	}
	//session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}else{
		//defer session.Close()

		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)
		return session
	}

}