package Req

import (
	"WebSocket/Common"
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//创建角色 1007
func CreateRoleReq(Uid int32,MsgId int32) []byte {
	Name := "kjcx" + Common.GetRandomString(3)
	CreateRoleReq := &AutoMsg.CreateRoleReq{
		Name: Name,
		Sex:  1,
	}
	data, _ := proto.Marshal(CreateRoleReq)
	Param,_ := json.Marshal(CreateRoleReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: MsgId,
		Name: "创建角色请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(MsgId, data)
}