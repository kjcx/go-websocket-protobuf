package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

type FriendSearch struct {
	Name string
	Search bool
}
//搜索好友 1024
func FriendSearchReq(FriendSearch *FriendSearch) []byte {
	FriendSearchReq := &AutoMsg.FriendSearchReq{
		Name:   FriendSearch.Name,
		Search: FriendSearch.Search,
	}
	Data,_ := proto.Marshal(FriendSearchReq)
	return SendRev.Send(1024,Data)
}
