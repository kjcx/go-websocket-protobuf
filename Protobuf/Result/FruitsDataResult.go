package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//水果机列表返回 1030
func FruitsDataResult(Uid int32,Data []byte) *AutoMsg.FruitsDataResult  {
	FruitsDataResult := &AutoMsg.FruitsDataResult{}
	proto.Unmarshal(Data,FruitsDataResult)
	Param,_ := json.Marshal(FruitsDataResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1030,
		Name:  "水果机列表返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return FruitsDataResult
}
