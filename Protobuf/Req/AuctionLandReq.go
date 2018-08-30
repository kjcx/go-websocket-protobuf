package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type AuctionLand struct {
	Pos int32
}
//2007 竞拍请求
func AuctionLandReq(Uid int32,AuctionLand *AuctionLand) []byte{
	AuctionLandReq := &AutoMsg.AuctionLandReq{}
	Data,_ := proto.Marshal(AuctionLandReq)
	Param,_ := json.Marshal(AuctionLandReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2007,
		Name:  "竞拍请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2007,Data)
}
