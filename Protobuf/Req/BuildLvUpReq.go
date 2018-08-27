package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)


type  BuildLvUp struct {
	BuildId []string
}
//升级店铺请求1018
func BuildLvUpReq(Uid int32,BuildLvUp *BuildLvUp) []byte{
	BuildLvUpReq := &AutoMsg.BuildLvUpReq{
		BuildId: BuildLvUp.BuildId,
	}
	Data,_ :=proto.Marshal(BuildLvUpReq)
	return SendRev.Send(1018,Data)
}
