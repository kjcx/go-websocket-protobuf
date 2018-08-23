package SendRev

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/proto/test_proto"
	"fmt"
	"WebSocket/Mgo"
	"strconv"
	"time"
)

//发送消息
func Send(MsgID int32, Data []byte) []byte {

	MsgSend := &AutoMsg.MsgBaseRev{
		MsgId: MsgID,
		Data:  Data,
	}
	data, _ := proto.Marshal(MsgSend)
	return data
}
//接收数据
func Rev(Data []byte) *AutoMsg.MsgBaseSend {
	fmt.Println("data:")
	MsgRev := &AutoMsg.MsgBaseSend{}
	err := proto.Unmarshal(Data, MsgRev)
	if err != nil {
		fmt.Println("marshaling error1111: ", err)
	}
	MsgID := MsgRev.GetMsgID()
	value :=  MsgRev.GetResult()
	if value > 0{
		mgo := Mgo.Mongo()
		result :=Mgo.WsResult(mgo,strconv.Itoa(int(value)))
		Name := Mgo.GetMsg(mgo,MsgID)
		log := &Mgo.Log{
			Uid:   14,
			MsgId: MsgID,
			Name:  Name,
			Param: "",
			Date:  time.Now(),
			Msg:   result.Msg,
		}
		mongo := Mgo.Mongo()
		Mgo.InsertLog(mongo,log)
		fmt.Println("有错误消息")

	}

	fmt.Println(MsgID)
	//MsgRev.GetData()
	return MsgRev
}