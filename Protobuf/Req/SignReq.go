package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)


//请求签到列表1121
func SignReq() ([]byte){
	Sign := &AutoMsg.SignReq{}
	data,_ := proto.Marshal(Sign)
	return SendRev.Send(1121,data)
}
//请求签到 1122
func DaySignReq(Day int32) ([]byte) {
	DaySignReq := &AutoMsg.DaySignReq{
		Day:Day,
	}
	data,_ := proto.Marshal(DaySignReq)
	return SendRev.Send(1122,data)
}

