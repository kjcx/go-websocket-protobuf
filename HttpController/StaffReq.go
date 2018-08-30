package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"fmt"

	"encoding/json"
	"WebSocket/Protobuf/Req"
	"WebSocket/Ws"
	"WebSocket/Mgo"
	"strings"
	"time"
)
//加载员工 1083
func LoadStaffReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	fmt.Println("请求开始时间",time.Now())
	param := Common.HttpParam{R:r,Ps:ps}
	str := Req.LoadStaffReq(param.GetParam().Uid)
	Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(time.Second * 1)
	//读取1118信息并返回
	result :=Mgo.FindOne(param.GetParam().Uid,1118)
	//w.Write([]byte("加载员工"))
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	//res,_ := json.Marshal(result)
	//w.Write(res)
	w.Write([]byte(result))
}
//招聘 1082
func RefStaffReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R:r,Ps:ps}
	data :=  &Req.RefStaff{}
	json.Unmarshal(param.GetParam().Body,data)
	fmt.Println(data)
	str := Req.RefStaffReq(param.GetParam().Uid,data)
	Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("招聘请求"))
}
//调入调出员工请求 1091
func ComeOutEmployeeReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

	param := Common.HttpParam{R:r,Ps:ps}
	data := &Req.ComeOutEmployee{}
	json.Unmarshal(param.GetParam().Body,data)
	fmt.Println(data)
	str := Req.ComeOutEmployeeReq(param.GetParam().Uid,data)
	Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("调入调出员工"))

}
//培训员工请求 1086
func CultivateEmployeeReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	fmt.Println("请求时间",time.Now())
	param := Common.HttpParam{R:r,Ps:ps}
	//r.ParseForm()
	//fmt.Println("form",r.Form.Get("ListId"))
	//fmt.Println("form",r.Form)
	//ListId := r.Form.Get("ListId")
	//fmt.Println(ListId)
	//str_arr := strings.Split(ListId,",")
	//fmt.Println(str_arr)
	//fmt.Println(string(param.GetParam().Body))
	data := &Req.CultivateEmployee{}
	json.Unmarshal(param.GetParam().Body,data)
	ListId := data.ListId
	str_arr := strings.Split(ListId[0],",")
	fmt.Println(str_arr)
	data.ListId = str_arr
	str := Req.CultivateEmployeeReq(param.GetParam().Uid,data)
	fmt.Println(data)
	Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	//w.Write([]byte("培训员工请求"))
	          //返回数据格式是json
	result :=Mgo.FindOne(param.GetParam().Uid,1121)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))



}
//解雇员工请求 1087
func DismissalEmployeeReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	data := &Req.DismissalEmployee{}
	json.Unmarshal(param.GetParam().Body,data)
	str := Req.DismissalEmployeeReq(param.GetParam().Uid,data)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	time.Sleep(2* time.Second)
	result:=Mgo.FindOne(param.GetParam().Uid,1122)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("解雇员工请求")

}

