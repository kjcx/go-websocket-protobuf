package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)
type ModelClothes struct {
	Id []int32
}
//购买时装 1150
func ModelClothesReq(ModelClothes *ModelClothes) []byte {
	ModelClothesReq := &AutoMsg.ModelClothesReq{
		Id: ModelClothes.Id,
	}
	Data,_ := proto.Marshal(ModelClothesReq)
	return SendRev.Send(1150,Data)
}
