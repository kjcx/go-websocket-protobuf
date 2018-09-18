package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"WebSocket/Mgo"
	"time"
	"encoding/json"
)

type GiveGift struct {
	ItemCount            map[int32]int32
	RoleId               int32
}
//赠送礼物 2038
func GiveGiftReq(Uid int32,GiveGift *GiveGift) []byte{
	GiveGiftReq := &AutoMsg.GiveGiftReq{
		ItemCount: GiveGift.ItemCount,
		RoleId:    GiveGift.RoleId,
	}

	Data,_ := proto.Marshal(GiveGiftReq)
	Param,_:= json.Marshal(GiveGiftReq)
	log := &Mgo.Log{
		Uid:   Uid,
		MsgId: 2037,
		Name:  "赠送礼物",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return SendRev.Send(2037,Data)

}