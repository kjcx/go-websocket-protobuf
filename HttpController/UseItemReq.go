package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Protobuf/Req"
	"WebSocket/Ws"
	"fmt"
	"WebSocket/Common"
	"encoding/json"
)

//使用道具返回
func UseItemReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	UseItem := &Req.UseItem{}
	json.Unmarshal(param.GetParam().Body,UseItem)
	Data := Req.UseItemReq(param.GetParam().Uid,UseItem)
	go Ws.ChanMsgWrite(Ws.SendChan{Uid: param.GetParam().Uid, Data: Data})
	w.Write([]byte("使用道具返回"))
	fmt.Println("使用道具返回")
}