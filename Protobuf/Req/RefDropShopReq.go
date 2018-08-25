package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)
type RefDropShop struct {
	ShopTypeId int32 // 商店id
}
//发送刷新商城商品请求 1077
func (R RefDropShop)RefDropShopReq() []byte{
	RefDropShopReq := &AutoMsg.RefDropShopReq{
		ShopTypeId: R.ShopTypeId,
	}
	Data, _ := proto.Marshal(RefDropShopReq)
	return SendRev.Send(1077,Data)
}
