package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"

	"encoding/json"
	"WebSocket/Protobuf/Req"
	"fmt"
	"WebSocket/Ws"
	"time"
	"WebSocket/Mgo"
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
	data := &Req.FriendApply{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.FriendApplyReq(param.GetParam().Uid,data)
	fmt.Println(str)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("申请加好友返回"))
}

//通过加好友 1022
func FriendAddReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.FriendAdd{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.FriendAddReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(1* time.Second)
	result := Mgo.FindOne(param.GetParam().Uid,1013)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("通过加好友")
}
//拒绝申请好友
func FriendApplyClearReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.FriendApplyClear{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.FriendApplyClearReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("拒绝申请好友"))
}
//搜索好友 1024
func FriendSearchReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.FriendSearch{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.FriendSearchReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	result := Mgo.FindOne(param.GetParam().Uid,1016)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("搜索好友")
}


//拉黑好友请求 1026
func FriendAddBlackReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.FriendAddBlack{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.FriendAddBlackReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,1018)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("拉黑好友请求")
}

//移除黑名单 1027
func FriendRemoveBlackReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.FriendRemoveBlack{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.FriendRemoveBlackReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,1019)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("移除黑名单请求")
}

