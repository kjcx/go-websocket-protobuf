package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//商战胜利收取道具返回2048
func CollectItemResult(Uid int32,Data []byte)*AutoMsg.CollectItemResult  {
	CollectItemResult := &AutoMsg.CollectItemResult{}
	proto.Unmarshal(Data,CollectItemResult)
	Param,_ := json.Marshal(CollectItemResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2048,
		Name:  "商战胜利收取道具返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return CollectItemResult
}
