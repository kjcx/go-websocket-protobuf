package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)
type FriendApply struct {
	RoleId string
}
//1021 申请加好友
func FriendApplyReq(FriendApply *FriendApply) []byte{
	FriendApplyReq := &AutoMsg.FriendApplyReq{
		RoleId: FriendApply.RoleId,
	}
	Data,_ := proto.Marshal(FriendApplyReq)
	return SendRev.Send(1021,Data)
}
