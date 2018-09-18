package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)
//心跳发送
func HeartbeatReq(Uid int32) []byte{
	HeartbeatReq := &AutoMsg.HeartbeatReq{}
	Data,_ := proto.Marshal(HeartbeatReq)
	return SendRev.Send(5024,Data)
}
