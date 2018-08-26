package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//批量申请好友返回 1013
func FriendAddResult(Uid int32,Data []byte) *AutoMsg.FriendAddResult{
	FriendAddResult := &AutoMsg.FriendAddResult{}
	proto.Unmarshal(Data,FriendAddResult)
	Param,_ := json.Marshal(FriendAddResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1013,
		Name:"批量申请好友返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return FriendAddResult
}
