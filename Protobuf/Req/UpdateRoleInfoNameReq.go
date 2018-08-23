package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
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
	return Send(1103,Data)
}
