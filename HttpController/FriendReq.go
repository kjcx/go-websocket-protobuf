package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"

	"encoding/json"
	"WebSocket/Protobuf/Req"
	"fmt"
	"WebSocket/Ws"
)

//赠送礼物请求返回 2038
func GiveGiftReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.GiveGift{}
	json.Unmarshal(param.GetParam().Body,data)
	fmt.Println(string(param.GetParam().Body))
	str := Req.GiveGiftReq(param.GetParam().Uid,data)
	fmt.Println(str)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("赠送礼物请求返回"))
}

//礼物列表请求返回 2056
func GiveListReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := Common.HttpParam{R: r, Ps: ps}
	fmt.Println(string(param.GetParam().Body))
	str := Req.GiveListReq(param.GetParam().Uid)
	fmt.Println(str)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("礼物列表请求返回"))
}
//申请加好友 1021
func FriendApplyReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := Common.HttpParam{R: r, Ps: ps}
	fmt.Println(string(param.GetParam().Body))
	data := &Req.FriendApply{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.FriendApplyReq(param.GetParam().Uid,data)
	fmt.Println(str)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("申请加好友返回"))
}

//申请加好友 1022
func FriendAddReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.FriendAdd{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.FriendAddReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("申请加好友返回"))
}

