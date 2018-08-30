package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)
//竞拍请求返回 2008
func AuctionLandResult(Uid int32,Data []byte)  *AutoMsg.AuctionLandResult {
	AuctionLandResult := &AutoMsg.AuctionLandResult{}
	proto.Unmarshal(Data,AuctionLandResult)
	Param,_ := json.Marshal(AuctionLandResult)
	log := Mgo.Log{
		Uid:   Uid,
		MsgId: 2008,
		Name:  "竞拍请求返回",
		Param: string(Param),
		Date:  time.Now(),
		Msg:   "无",
	}
	log.InsertLog()
	return AuctionLandResult
}
