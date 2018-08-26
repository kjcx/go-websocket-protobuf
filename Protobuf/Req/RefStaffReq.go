package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type RefStaff struct {
	TypeId int32
}
//1082 招聘
func RefStaffReq(Uid int32,RefStaff *RefStaff) []byte {
	RefStaffReq := &AutoMsg.RefStaffReq{
		TypeId: RefStaff.TypeId,
	}
	Data,_ := proto.Marshal(RefStaffReq)
	Param,_ := json.Marshal(RefStaffReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1082,
		Name: "招聘请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1082,Data)
}
