package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//谈判团居民任命返回 2032
func TalkingGroupResidentResult(Uid int32,Data []byte) *AutoMsg.TalkingGroupResidentResult  {
	TalkingGroupResidentResult := &AutoMsg.TalkingGroupResidentResult{}
	proto.Unmarshal(Data,TalkingGroupResidentResult)
	Param,_ := json.Marshal(TalkingGroupResidentResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2032,
		Name:  "谈判团居民任命返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return TalkingGroupResidentResult
}
