package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"WebSocket/Protobuf/Req"
	"WebSocket/Ws"
	"WebSocket/Mgo"
	"fmt"
	"encoding/json"
	"time"
)

//水果机列表 1031
func FruitsDataReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	str := Req.FruitsDataReq(param.GetParam().Uid)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	result:=Mgo.FindOne(param.GetParam().Uid,1030)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("水果机列表请求")

}

//水果机请求 1035
func RaffleFruitsReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.RaffleFruits{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.RaffleFruitsReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,1035)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("水果机请求")

}
