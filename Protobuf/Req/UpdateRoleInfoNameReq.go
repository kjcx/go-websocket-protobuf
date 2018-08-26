package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type UpdateRoleInfoName struct {
	RoleName string
}
//修改角色名称 1103
func UpdateRoleInfoNameReq(Uid int32,UpdateRoleInfoName *UpdateRoleInfoName) []byte {
	UpdateName := &AutoMsg.UpdateRoleInfoNameReq{
		RoleName:UpdateRoleInfoName.RoleName,
	}
	Data,_ := proto.Marshal(UpdateName)
	Param,_ := json.Marshal(UpdateName)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1103,
		Name: "修改角色名称请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1103,Data)
}
