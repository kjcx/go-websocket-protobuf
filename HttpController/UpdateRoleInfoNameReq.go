package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Protobuf/Req"
	"fmt"
	"WebSocket/Common"
	"encoding/json"
	"WebSocket/Ws"
)

func UpdateRoleInfoNameReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := Common.HttpParam{R: r, Ps: ps}
	UpdateRoleInfoName := &Req.UpdateRoleInfoName{}
	json.Unmarshal(param.GetParam().Body,UpdateRoleInfoName)
	str := Req.UpdateRoleInfoNameReq(param.GetParam().Uid,UpdateRoleInfoName)
	fmt.Println(str)
	Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	//arr := Ws.GetWebSocket()
	//go Ws.Send(arr[14],str)
}
