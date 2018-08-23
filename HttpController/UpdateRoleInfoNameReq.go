package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Protobuf/Req"
	"fmt"
	"WebSocket/Ws"
)

func UpdateRoleInfoNameReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	RoleName := ps.ByName("name")
	str := Req.UpdateRoleInfoNameReq(&Req.UpdateRoleInfoName{RoleName: RoleName})
	fmt.Println(str)
	arr := Ws.GetWebSocket()
	go Ws.Send(arr[3000],str)
}
