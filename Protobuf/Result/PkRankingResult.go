package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//pk排行榜返回 2028
func PkRankingResult(Uid int32,Data []byte) *AutoMsg.PkRankingResult {
	PkRankingResult := &AutoMsg.PkRankingResult{}
	proto.Unmarshal(Data,PkRankingResult)
	Param,_ := json.Marshal(PkRankingResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2028,
		Name:  "pk排行榜返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return PkRankingResult
}
