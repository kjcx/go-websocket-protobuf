package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

type FriendRemove struct {
	RoleId string
}
//删除好友请求 1023
func FriendRemoveReq(FriendRemove *FriendRemove) []byte{
	FriendRemoveReq := &AutoMsg.FriendRemoveReq{
		RoleId: FriendRemove.RoleId,
	}
	Data,_ := proto.Marshal(FriendRemoveReq)
	return SendRev.Send(1023,Data)
}
