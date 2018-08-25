package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

type FriendAdd struct {
	RoleIds []string
}
//1022 批量申请好友
func FriendAddReq(FriendAdd *FriendAdd) []byte{
	FriendAddReq := &AutoMsg.FriendAddReq{
		RoleIds: FriendAdd.RoleIds,
	}
	Data,_ := proto.Marshal(FriendAddReq)
	return SendRev.Send(1022,Data)
}
