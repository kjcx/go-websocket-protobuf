package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

type UpdateRoleInfoName struct {
	RoleName string
}
//修改角色名称 1103
func UpdateRoleInfoNameReq(UpdateRoleInfoName *UpdateRoleInfoName) []byte {
	UpdateName := &AutoMsg.UpdateRoleInfoNameReq{
		RoleName:UpdateRoleInfoName.RoleName,
	}
	Data,_ := proto.Marshal(UpdateName)
	return SendRev.Send(1103,Data)
}
