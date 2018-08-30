package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)
//pk排行榜 2027
func PkRankingReq(Uid int32)  []byte{
	PkRankingReq := &AutoMsg.PkRankingReq{}
	Data,_ := proto.Marshal(PkRankingReq)
	Param,_ := json.Marshal(PkRankingReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2027,
		Name:  "pk排行榜",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2027,Data)
}
