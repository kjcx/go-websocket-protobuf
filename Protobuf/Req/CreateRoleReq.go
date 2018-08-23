package Req

import (
	"WebSocket/Common"
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
)

//创建角色 1007
func CreateRoleReq(MsgId int32) []byte {
	Name := "kjcx" + Common.GetRandomString(3)
	CreateRoleReq := &AutoMsg.CreateRoleReq{
		Name: Name,
		Sex:  1,
	}
	data, _ := proto.Marshal(CreateRoleReq)
	return Send(MsgId, data)
}