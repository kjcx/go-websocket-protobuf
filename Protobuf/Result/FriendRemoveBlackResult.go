package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//移除黑名单返回 1019
func FriendRemoveBlackResult(Uid int32,Data []byte)*AutoMsg.FriendRemoveBlackResult{
	FriendRemoveBlackResult := &AutoMsg.FriendRemoveBlackResult{}
	proto.Unmarshal(Data,FriendRemoveBlackResult)
	Param,_ := json.Marshal(FriendRemoveBlackResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1019,
		Name:  "移除黑名单返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return FriendRemoveBlackResult

}
