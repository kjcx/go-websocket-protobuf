package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//1064 请求竞拍土地返回
func GetMapResult(Uid int32,Data []byte) *AutoMsg.GetMapResult{
	GetMapResult := &AutoMsg.GetMapResult{}
	proto.Unmarshal(Data,GetMapResult)
	Param,_ := json.Marshal(GetMapResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1064,
		Name:  "请求竞拍土地返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return GetMapResult
}
