package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type GetAuctionLand struct {
	Zone string
}
// 今日竞拍列表请求 2005
func GetAuctionLandReq(Uid int32,GetAuctionLand *GetAuctionLand) []byte{
	GetAuctionLandReq := &AutoMsg.GetAuctionLandReq{
		Zone: GetAuctionLand.Zone,
	}
	Data,_ := proto.Marshal(GetAuctionLandReq)
	Param,_ := json.Marshal(GetAuctionLandReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2005,
		Name:  "今日竞拍列表请求",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2005,Data)
}
