package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Protobuf/Req"
	"fmt"
	"WebSocket/Ws"
	"encoding/json"
	"WebSocket/Common"
	"WebSocket/Mgo"
)


//提升品质返回
func NpcRelationAdvanceReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.NpcRelationAdvance{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.NpcRelationAdvanceReq(param.GetParam().Uid,data)
	fmt.Println(str)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("提升品质返回"))
}

//提升好感度
func AddNpcFavorabilityReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	favorability := &Req.AddNpcFavorability{}

	json.Unmarshal(param.GetParam().Body,favorability)
	Data := Req.AddNpcRelationAdvanceReq(param.GetParam().Uid,favorability)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: Data})
	w.Write([]byte("AddNpcFavorability提升好感度"))

	fmt.Println("提升好感度")

}
//解锁npc请求
func UnlockNpcReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	UnlockNpc := &Req.UnlockNpc{}
	json.Unmarshal(param.GetParam().Body,UnlockNpc)
	str := Req.UnlockNpcReq(param.GetParam().Uid,UnlockNpc)
	Ws.ChanMsgWrite(Ws.SendChan{Uid:param.GetParam().Uid, Data: str})
	w.Write([]byte("解锁npc请求"))
	fmt.Println("解锁npc请求")

}

//加载npc好感度
func NpcFavorabilityReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}

	str := Req.NpcFavorabilityReq(param.GetParam().Uid)
	Ws.ChanMsgWrite(Ws.SendChan{Uid:param.GetParam().Uid, Data: str})
	w.Write([]byte("加载npc请求"))
	fmt.Println("加载npc请求")

}

//居民委托任务请求 1096
func ResidentDelegateReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	str := Req.ResidentDelegateReq(param.GetParam().Uid)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	result:=Mgo.FindOne(param.GetParam().Uid,1134)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("居民委托任务请求")

}
