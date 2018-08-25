package Req

import (
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/golang/protobuf/proto/test_proto"
	"WebSocket/Protobuf"
)


//结构体
type UseItem struct {
	ItemId int32
	Count int32
}
//使用道具 1015
func UseItemReq(UseItem *UseItem) []byte{
	//data := &UseItem{10010,1}
	UseItemReq := &AutoMsg.UseItemReq{
		ItemId:UseItem.ItemId,
		Count:UseItem.Count,
	}
	Data,err := proto.Marshal(UseItemReq)
	if err!=nil{
		log.Fatal("UseItemReq解析失败")
	}
	return SendRev.Send(1015,Data)
}
