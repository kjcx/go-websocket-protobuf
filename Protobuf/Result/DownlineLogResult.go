package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//下线带 2042来奖励日志返回
func DownlineLogResult(Uid int32,Data []byte) *AutoMsg.DownlineLogResult {
	DownlineLogResult := &AutoMsg.DownlineLogResult{}
	proto.Unmarshal(Data,DownlineLogResult)
	Param,_ := json.Marshal(DownlineLogResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2042,
		Name:  "下线带来奖励日志返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return DownlineLogResult
}
