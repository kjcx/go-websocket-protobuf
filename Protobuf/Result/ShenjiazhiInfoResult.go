package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//身价值明细返回 2050
func ShenjiazhiInfoResult(Uid int32,Data[]byte) *AutoMsg.ShenjiazhiInfoResult {
	ShenjiazhiInfoResult:= &AutoMsg.ShenjiazhiInfoResult{}
	proto.Unmarshal(Data,ShenjiazhiInfoResult)
	Param,_ := json.Marshal(ShenjiazhiInfoResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2050,
		Name:  "身价值明细返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return ShenjiazhiInfoResult
}
