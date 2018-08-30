package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
	"WebSocket/Protobuf"
)

//自己竞拍列表 2009
func MyLandInfoReq(Uid int32) []byte{
	MyLandInfoReq := &AutoMsg.GetMyLandInfoReq{}
	Data,_ := proto.Marshal(MyLandInfoReq)
	Param,_ := json.Marshal(MyLandInfoReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2009,
		Name:  "自己竞拍列表",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2009,Data)
}
