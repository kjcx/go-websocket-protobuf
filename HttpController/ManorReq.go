package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"WebSocket/Protobuf/Req"
	"fmt"
	"WebSocket/Ws"
	"encoding/json"
	"time"
	"WebSocket/Mgo"
	"github.com/golang/protobuf/proto/test_proto"
)


//请求庄园
func RequestManorReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.RequestManor{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.RequestManorReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("请求庄园"))
}
//拜访请求
func ManorVisitInfoReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	str := Req.ManorVisitInfoReq(param.GetParam().Uid)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("拜访请求"))
}


//偷菜返回 1061
func StealSemenReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &AutoMsg.StealSemenReq{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.StealSemenReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,1091)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("偷菜返回")
}


//随机请求庄园 2013
func RandManorReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	str := Req.RandManorReq(param.GetParam().Uid)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,2014)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("随机请求庄园")
}

