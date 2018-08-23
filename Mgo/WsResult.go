package Mgo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)



type Result struct {
	Type string
	Msg string
	Value string
}

func WsResult(mgo *mgo.Session,value string) Result  {
	c := mgo.DB("ckzc").C("WsResult")
	result := Result{}
	c.Find(bson.M{"value": value}).One(&result)
	fmt.Println(result.Msg)
	return result
}