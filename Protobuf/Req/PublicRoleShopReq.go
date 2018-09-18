package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)
//1141 自己公有区店铺请求
func PublicRoleShopReq(Uid int32) []byte{
	PublicRoleShopReq := &AutoMsg.PublicRoleShopReq{}
	Data,_ := proto.Marshal(PublicRoleShopReq)
	return SendRev.Send(1141,Data)
}
