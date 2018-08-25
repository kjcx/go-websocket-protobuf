package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

//1025
type FriendApplyClear struct {
	RoleIds []string
}
//拒绝用户申请 1025
func FriendApplyClearReq(FriendApplyClear *FriendApplyClear) []byte {
	FriendApplyClearReq := &AutoMsg.FriendApplyClearReq{
		RoleIds: FriendApplyClear.RoleIds,
	}
	Data,_ := proto.Marshal(FriendApplyClearReq)
	return SendRev.Send(1025,Data)
}
