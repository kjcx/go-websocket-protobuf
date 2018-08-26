package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//删除好友返回1012
func FriendRemoveResult(Uid int32,Data []byte) *AutoMsg.FriendRemoveResult{
	FriendRemoveResult := &AutoMsg.FriendRemoveResult{}
	proto.Unmarshal(Data,FriendRemoveResult)
	Param,_ := json.Marshal(FriendRemoveResult)
	log :=&Mgo.Log{
		Uid:Uid,
		MsgId:1012,
		Name:"删除好友返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return FriendRemoveResult
}
