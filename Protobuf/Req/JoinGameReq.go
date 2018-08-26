package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"WebSocket/Protobuf"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//加入游戏 1012
func JoinGameReq(Uid int32,MsgId int32) []byte {
	JoinGameReq := &AutoMsg.JoinGameReq{}
	data, _ := proto.Marshal(JoinGameReq)
	Param,_ := json.Marshal(JoinGameReq)
	//插入日志
	log := &Mgo.Log{
		Uid:Uid,
		MsgId: MsgId,
		Name: "加入游戏请求",
		Param: string(Param),
		Date: time.Now(),
		Msg: "",
	}
	log.InsertLog()
	return SendRev.Send(MsgId, data)
}