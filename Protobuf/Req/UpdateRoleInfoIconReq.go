package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
type UpdateRoleInfoIcon struct {
	RoleIcon string
}
//修改头像 1102
func UpdateRoleInfoIconReq(Uid int32,UpdateRoleInfoIcon *UpdateRoleInfoIcon) []byte{
	UpdateRoleInfoIconReq := &AutoMsg.UpdateRoleInfoIconReq{
		RoleIcon: UpdateRoleInfoIcon.RoleIcon,
	}
	Data,_ := proto.Marshal(UpdateRoleInfoIconReq)
	Param,_ := json.Marshal(UpdateRoleInfoIconReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1102,
		Name: "修改头像请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1102,Data)
}
