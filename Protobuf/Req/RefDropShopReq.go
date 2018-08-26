package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
type RefDropShop struct {
	ShopTypeId int32 // 商店id
}
//发送刷新商城商品请求 1077
func (R RefDropShop)RefDropShopReq(Uid int32) []byte{
	RefDropShopReq := &AutoMsg.RefDropShopReq{
		ShopTypeId: R.ShopTypeId,
	}
	Data, _ := proto.Marshal(RefDropShopReq)
	Param,_ := json.Marshal(RefDropShopReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1077,
		Name: "发送刷新商城商品请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1077,Data)
}
