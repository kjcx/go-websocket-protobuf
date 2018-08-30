package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"WebSocket/Protobuf/Req"
	"encoding/json"
	"fmt"
	"WebSocket/Ws"
	"WebSocket/Mgo"
	"time"
)
//人才市场列表请求
func GetTalentListReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	str := Req.GetTalentListReq(param.GetParam().Uid)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("人才市场列表请求"))
}
//升级店铺请求 1018
func BuildLvUpReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.BuildLvUp{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.BuildLvUpReq(param.GetParam().Uid,data)
	fmt.Println(data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(1*time.Second)
	result := Mgo.FindOne(param.GetParam().Uid,1004)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("升级店铺请求")
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
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("刷新人才市场"))
}


//创建店铺 1005
func CreateBuildReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.CreateBuild{}
	str := Req.CreateBuildReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	//1058
	result:=Mgo.FindOne(param.GetParam().Uid,1058)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("创建店铺")

}

//创建店铺 1006
func CreateCompanyReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.CreateCompany{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.CreateCompanyReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	//1058
	result:=Mgo.FindOne(param.GetParam().Uid,1059)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("创建公司返回")

}

//今日竞拍列表请求 2005
func GetAuctionLandReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.GetAuctionLand{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.GetAuctionLandReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	//1058
	result:=Mgo.FindOne(param.GetParam().Uid,2006)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("今日竞拍列表请求")

}
//竞拍请求 2007
func AuctionLandReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.AuctionLand{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.AuctionLandReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	//1058
	result:=Mgo.FindOne(param.GetParam().Uid,2008)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("竞拍请求")

}
//自己竞拍列表2009
func MyLandInfoReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}

	str := Req.MyLandInfoReq(param.GetParam().Uid)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	result:=Mgo.FindOne(param.GetParam().Uid,2010)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("自己竞拍列表请求")

}



