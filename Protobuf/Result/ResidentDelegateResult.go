package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//居民委托任务返回 1134
func ResidentDelegateResult(Uid int32,Data []byte) *AutoMsg.ResidentDelegateResult {
	ResidentDelegateResult := &AutoMsg.ResidentDelegateResult{}
	proto.Unmarshal(Data,ResidentDelegateResult)
	Param,_ := json.Marshal(ResidentDelegateResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1134,
		Name:  "居民委托任务返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return ResidentDelegateResult
}
