package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
)


//请求签到列表
func SignReq() ([]byte){
	Sign := &AutoMsg.SignReq{}
	data,_ := proto.Marshal(Sign)
	return Send(1121,data)
}
//请求签到
func DaySignReq(Day int32) ([]byte) {
	DaySignReq := &AutoMsg.DaySignReq{
		Day:Day,
	}
	data,_ := proto.Marshal(DaySignReq)
	return Send(1122,data)
}


//发送消息
func Send(MsgID int32, Data []byte) []byte {

	MsgSend := &AutoMsg.MsgBaseRev{
		MsgId: MsgID,
		Data:  Data,
	}
	data, _ := proto.Marshal(MsgSend)
	return data
}