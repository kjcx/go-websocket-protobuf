package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//水果机请求 1035
func RaffleFruitsResult(Uid int32,Data []byte) *AutoMsg.RaffleFruitsResult{
	RaffleFruitsResult := &AutoMsg.RaffleFruitsResult{}
	proto.Unmarshal(Data,RaffleFruitsResult)
	Param,_ := json.Marshal(RaffleFruitsResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1035,
		Name:  "水果机请求返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return RaffleFruitsResult
}
