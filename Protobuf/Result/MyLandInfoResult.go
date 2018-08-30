package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//自己竞拍列表返回 2010
func MyLandInfoResult(Uid int32,Data []byte) *AutoMsg.MyLandInfoResult{
	MyLandInfoResult := &AutoMsg.MyLandInfoResult{}
	proto.Unmarshal(Data,MyLandInfoResult)
	Param,_ := json.Marshal(MyLandInfoResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2010,
		Name:  "自己竞拍列表返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return MyLandInfoResult
}
