package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"WebSocket/Protobuf/Req"
	"encoding/json"
	"fmt"
	"WebSocket/Ws"
)
//人才市场列表请求
func GetTalentListReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	str := Req.GetTalentListReq(param.GetParam().Uid)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("人才市场列表请求"))
}
//升级店铺请求
func BuildLvUpReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.BuildLvUp{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.BuildLvUpReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("升级店铺请求"))
}
//雇佣经理请求
func TalentHireReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.TalentHire{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.TalentHireReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("雇佣经理请求"))
}
//解雇经理请求
func TalentFireReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.TalentFire{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.TalentFireReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("解雇经理请求"))
}
//请求竞拍土地
func GetMapReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.GetMap{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.GetMapReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("请求竞拍土地"))
}


//刷新人才市场
func TalentRefreshReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	str := Req.TalentRefreshReq(param.GetParam().Uid)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("刷新人才市场"))
}

