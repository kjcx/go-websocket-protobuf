package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

type ChangeAvatar struct {
	Id []int32
}
//改变时装请求
func ChangeAvatarReq(ChangeAvatar *ChangeAvatar) []byte {
	ChangeAvatarReq := &AutoMsg.ChangeAvatarReq{
		Id: ChangeAvatar.Id,
	}
	Data,_ := proto.Marshal(ChangeAvatarReq)
	return SendRev.Send(1002,Data)
}
