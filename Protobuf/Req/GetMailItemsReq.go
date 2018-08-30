package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"encoding/json"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
	"github.com/golang/protobuf/proto"
)

type GetMailItems struct {
	Id []string
}
//一键获取邮件奖品
func GetMailItemsReq(Uid int32,GetMailItems *GetMailItems) []byte{
	GetMailItemsReq := &AutoMsg.GetMailItemsReq{
		Id: GetMailItems.Id,
	}
	Data,_ := proto.Marshal(GetMailItemsReq)
	Param,_ := json.Marshal(GetMailItemsReq)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1065,
		Name:  "一键获取邮件物品",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(1065,Data)

}
