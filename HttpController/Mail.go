package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"WebSocket/Protobuf/Req"
	"encoding/json"
	"WebSocket/Ws"
	"time"
	"WebSocket/Mgo"
	"fmt"
)


//阅读邮件请求 1081
func ReadMailReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.ReadMail{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.ReadMailReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,1114)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("阅读邮件请求")

}

//一键获取邮件物品请求 1065
func GetMailItemsReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.GetMailItems{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.GetMailItemsReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,1095)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("一键获取邮件物品请求")

}
//批量删除邮件返回 1080
func DelMailsReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.DelMails{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.DelMailsReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,1113)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("批量删除邮件返回")

}
