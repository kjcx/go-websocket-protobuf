package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
	"encoding/json"
	"WebSocket/Mgo"
	"time"
)

//解析商店 1145
func ShopAllResult(Uid int32,Data []byte) *AutoMsg.ShopAllResult {
	ShopAllResult := &AutoMsg.ShopAllResult{}
	err := proto.Unmarshal(Data, ShopAllResult)
	if err != nil {
		fmt.Println("解析ShopAllResult出错")
	}
	Param,_ := json.Marshal(ShopAllResult)
	log := &Mgo.Log{
		Uid:Uid,
		MsgId:1168,
		Name:"签到列表",
		Param:string(Param),
		Date:time.Now(),
		Msg:"无",
	}
	log.InsertLog()
	return ShopAllResult
}
