package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

type FriendSearch struct {
	Name string
	Search bool
}
//搜索好友 1024
func FriendSearchReq(Uid int32,FriendSearch *FriendSearch) []byte {
	FriendSearchReq := &AutoMsg.FriendSearchReq{
		Name:   FriendSearch.Name,
		Search: FriendSearch.Search,
	}
	Data,_ := proto.Marshal(FriendSearchReq)
	Param,_ := json.Marshal(FriendSearchReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: 1024,
		Name: "搜索好友请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(1024,Data)
}
