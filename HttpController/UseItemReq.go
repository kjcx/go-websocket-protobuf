package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Protobuf/Req"
	"WebSocket/Ws"
	"fmt"
)

//使用道具返回
func UseItemReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	UseItem := &Req.UseItem{
		ItemId: 10001,
		Count:  1,
	}
	Data := Req.UseItemReq(UseItem)
	go Ws.ChanMsgWrite(Data)
	w.Write([]byte("使用道具返回"))

	fmt.Println("使用道具返回")

}