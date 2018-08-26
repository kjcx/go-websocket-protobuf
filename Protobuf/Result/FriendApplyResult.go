package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//1011 申请好友返回
func FriendApplyResult(Uid int32,Data []byte) *AutoMsg.FriendApplyResult{
	FriendApplyResult := &AutoMsg.FriendApplyResult{}
	proto.Unmarshal(Data,FriendApplyResult)
	Param,_ := json.Marshal(FriendApplyResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1011,
		Name:"申请好友返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return FriendApplyResult
}
