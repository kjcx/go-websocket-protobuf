package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
type ModelClothes struct {
	Id []int32
}
//购买时装 1150
func ModelClothesReq(Uid int32,ModelClothes *ModelClothes) []byte {
	ModelClothesReq := &AutoMsg.ModelClothesReq{
		Id: ModelClothes.Id,
	}
	Data,_ := proto.Marshal(ModelClothesReq)
	Param,_ := json.Marshal(LoadStaffReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1150,
		Name: "购买时装请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1150,Data)
}
