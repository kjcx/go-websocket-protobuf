package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//搜索好友返回 1016
func FriendSearchResult(Uid int32,Data []byte) *AutoMsg.FriendSearchResult{
	FriendSearchResult:= &AutoMsg.FriendSearchResult{}
	proto.Unmarshal(Data,FriendSearchResult)
	Param,_ := json.Marshal(FriendSearchResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1016,
		Name:"搜索好友返回",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return FriendSearchResult
}
