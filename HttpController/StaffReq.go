package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"fmt"

	"encoding/json"
	"WebSocket/Protobuf/Req"
	"WebSocket/Ws"
)
//加载员工 1083
func LoadStaffReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R:r,Ps:ps}
	str := Req.LoadStaffReq(param.GetParam().Uid)
	Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("加载员工"))
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
	str := Req.ComeOutEmployeeReq(param.GetParam().Uid,data)
	Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("调入调出员工"))

}
