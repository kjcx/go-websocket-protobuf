package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)


//请求签到列表1121
func SignReq(Uid int32) ([]byte){
	Sign := &AutoMsg.SignReq{}
	data,_ := proto.Marshal(Sign)
	Param,_ := json.Marshal(LoadStaffReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1121,
		Name: "签到列表请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1121,data)
}
//请求签到 1122
func DaySignReq(Uid int32,Day int32) ([]byte) {
	DaySignReq := &AutoMsg.DaySignReq{
		Day:Day,
	}
	data,_ := proto.Marshal(DaySignReq)
	Param,_ := json.Marshal(DaySignReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1122,
		Name: "签到请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1122,data)
}

