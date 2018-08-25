package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
)

//加入游戏 1012
func JoinGameReq(MsgId int32) []byte {
	JoinGameReq := &AutoMsg.JoinGameReq{}
	data, _ := proto.Marshal(JoinGameReq)
	return SendRev.Send(MsgId, data)
}