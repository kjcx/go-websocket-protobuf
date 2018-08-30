package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//施肥返回2012
func UseCompostResult(Uid int32,Data []byte) *AutoMsg.UseCompostResult {
	UseCompostResult :=&AutoMsg.UseCompostResult{}
	proto.Unmarshal(Data,UseCompostResult)
	Param,_ := json.Marshal(UseCompostResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2012,
		Name:  "施肥返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return UseCompostResult
}
