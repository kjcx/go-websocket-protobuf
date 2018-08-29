package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"WebSocket/Protobuf/Req"
	"fmt"
	"WebSocket/Ws"
	"encoding/json"
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
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("拜访请求"))
}
