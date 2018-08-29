package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//拜访记录返回 1201
func ManorVisitInfoResult(Uid int32,Data []byte) *AutoMsg.ManorVisitInfoResult{
	ManorVisitInfoResult := &AutoMsg.ManorVisitInfoResult{}
	proto.Unmarshal(Data,ManorVisitInfoResult)
	Param,_ := json.Marshal(ManorVisitInfoResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1201,
		Name:  "拜访记录返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return  ManorVisitInfoResult
}
