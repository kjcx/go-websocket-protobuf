package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Protobuf/Req"
	"fmt"
	. "WebSocket/Common"
	"encoding/json"
	"WebSocket/Ws"
)
//创建公司
func CreateBuildReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := HttpParam{R: r, Ps: ps}
	CreateBuild := &Req.CreateBuild{}
	json.Unmarshal(param.GetParam().Body,CreateBuild)
	str := Req.CreateBuildReq(param.GetParam().Uid,CreateBuild)
	Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: str})
	w.Write([]byte("创建公司返回"))
	fmt.Println("创建公司返回")
}
