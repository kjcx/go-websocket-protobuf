package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)
type SellItem struct {
	ItemId int32
	Count int64
}
//出售道具 1014
func SellItemReq(SellItem *SellItem) []byte{
	SellItemReq := &AutoMsg.SellItemReq{
		ItemId: SellItem.ItemId,
		Count:  SellItem.Count,
	}
	Data,_ := proto.Marshal(SellItemReq)
	return SendRev.Send(1014,Data)
}
