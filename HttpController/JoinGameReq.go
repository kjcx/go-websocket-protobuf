package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Common"
	"WebSocket/Protobuf/Req"
	"WebSocket/Ws"
	"fmt"
	"WebSocket/Mgo"
)

//进入游戏请求
func JoinGameReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	param := Common.HttpParam{R: r, Ps: ps}
	str := Req.JoinGameReq(param.GetParam().Uid)
	Ws.ChanMsgWrite(Ws.SendChan{Uid:param.GetParam().Uid, Data: str})
	result := Mgo.FindOne(param.GetParam().Uid,1066)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(result))
	fmt.Println("进入游戏请求")

}