package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//1182 刷新人才市场返回
func TalentRefreshResult(Uid int32,Data []byte) *AutoMsg.TalentRefreshResult{
	TalentRefreshResult := &AutoMsg.TalentRefreshResult{}
	proto.Unmarshal(Data,TalentRefreshResult)
	Param,_ := json.Marshal(TalentRefreshResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 1182,
		Name:  "刷新人才市场返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return TalentRefreshResult
}
