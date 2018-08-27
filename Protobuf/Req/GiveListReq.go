package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)
//礼物列表 2055
func GiveListReq(Uid int32) []byte{
	GiveListReq := &AutoMsg.GiveListReq{}
	Data,_ := proto.Marshal(GiveListReq)
	return SendRev.Send(2055,Data)

}
