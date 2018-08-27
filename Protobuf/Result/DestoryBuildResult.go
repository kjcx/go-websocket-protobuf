package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//出售店铺返回 1063
func DestoryBuildResult(Uid int32,Data []byte) *AutoMsg.DestoryBuildResult{
	DestoryBuildResult := &AutoMsg.DestoryBuildResult{}
	proto.Unmarshal(Data,DestoryBuildResult)
	Param,_ := json.Marshal(DestoryBuildResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1063,
		Name:  "出售店铺返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return DestoryBuildResult
}
