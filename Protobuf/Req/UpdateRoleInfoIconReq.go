package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)
type UpdateRoleInfoIcon struct {
	RoleIcon string
}
//修改头像 1102
func UpdateRoleInfoIconReq(UpdateRoleInfoIcon *UpdateRoleInfoIcon) []byte{
	UpdateRoleInfoIconReq := &AutoMsg.UpdateRoleInfoIconReq{
		RoleIcon: UpdateRoleInfoIcon.RoleIcon,
	}
	Data,_ := proto.Marshal(UpdateRoleInfoIconReq)
	return SendRev.Send(1102,Data)
}
