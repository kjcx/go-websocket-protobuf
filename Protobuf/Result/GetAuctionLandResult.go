package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
// 今日竞拍列表返回 2006
func GetAuctionLandResult(Uid int32,Data []byte) *AutoMsg.GetAuctionLandResult{
	GetAuctionLandResult :=&AutoMsg.GetAuctionLandResult{}
	proto.Unmarshal(Data,GetAuctionLandResult)
	Param,_ := json.Marshal(GetAuctionLandResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2006,
		Name:"今日竞拍列表返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return GetAuctionLandResult
}
