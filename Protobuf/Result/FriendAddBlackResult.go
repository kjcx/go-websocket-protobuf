package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//拉黑好友返回 1018
func FriendAddBlackResult(Uid int32,Data []byte) *AutoMsg.FriendAddBlackResult  {
	FriendAddBlackResult  := &AutoMsg.FriendAddBlackResult{}
	proto.Marshal(FriendAddBlackResult)
	Param,_ := json.Marshal(FriendAddBlackResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1018,
		Name:  "拉黑好友返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return  FriendAddBlackResult
}
