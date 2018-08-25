package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//拒绝申请返回 1017
func FriendApplyClearResult(Uid int32,Data []byte) {
	FriendApplyClearResult := &AutoMsg.FriendApplyClearResult{}
	proto.Unmarshal(Data,FriendApplyClearResult)
	Param,_ := json.Marshal(FriendApplyClearResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1017,
		Name:"拒绝申请返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	mongo := Mgo.Mongo()
	Mgo.InsertLog(mongo,log)
}
