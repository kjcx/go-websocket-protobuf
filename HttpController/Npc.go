package HttpController

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"WebSocket/Protobuf/Req"
	"fmt"
	"WebSocket/Ws"
)



func NpcRelationAdvanceReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := &Req.NpcRelationAdvance{
		NpcId:104,
	}
	str := Req.NpcRelationAdvanceReq(data)
	fmt.Println(str)
	go Ws.ChanMsgWrite(str)
	w.Write([]byte("sss"))
	w.Write([]byte("hhhh"))

	fmt.Println("not found")
}

//提升好感度
func AddNpcFavorabilityReq(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	favorability := &Req.AddNpcFavorability{
		104,
		105,
		100,
	}
	Data := Req.AddNpcRelationAdvanceReq(favorability)
	go Ws.ChanMsgWrite(Data)
	w.Write([]byte("AddNpcFavorability"))

	fmt.Println("提升好感度")

}