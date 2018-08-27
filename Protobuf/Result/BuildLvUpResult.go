package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//升级店铺返回1004
func BuildLvUpResult(Uid int32,Data []byte) *AutoMsg.BuildLvUpResult{
	BuildLvUpResult := &AutoMsg.BuildLvUpResult{}
	proto.Unmarshal(Data,BuildLvUpResult)
	Param,_ := json.Marshal(BuildLvUpResult)
	log := &Mgo.Log{
		Uid:   Uid,
		MsgId:     1004,
		Name:  "升级店铺返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return BuildLvUpResult
}
