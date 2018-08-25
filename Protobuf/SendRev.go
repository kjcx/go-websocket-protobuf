package SendRev

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/proto/test_proto"
	"fmt"
	"WebSocket/Mgo"
	"time"
	"strconv"
)

//发送消息
func Send(MsgID int32, Data []byte) []byte {
	MsgSend := &AutoMsg.MsgBaseRev{
		MsgId: MsgID,
		Data:  Data,
	}
	data, _ := proto.Marshal(MsgSend)
	mongo := Mgo.Mongo()
	Name := Mgo.GetMsg(mongo,MsgID)
	//插入日志
	log := &Mgo.Log{
		Uid:14,
		MsgId: MsgID,
		Name: Name,
		Param: "",
		Date: time.Now(),
		Msg: "",
	}
	Mgo.InsertLog(mongo,log)
	return data
}
//接收数据
func Rev(Uid int32,Data []byte) *AutoMsg.MsgBaseSend {
	fmt.Println("Uid",Uid,"data:",Data)
	MsgRev := &AutoMsg.MsgBaseSend{}
	err := proto.Unmarshal(Data, MsgRev)
	if err != nil {
		fmt.Println("marshaling error Rev: ", err)
	}
	MsgID := MsgRev.GetMsgID()
	value :=  MsgRev.GetResult()

	if value > 0{
		mongo := Mgo.Mongo()
		result :=Mgo.WsResult(mongo,strconv.Itoa(int(value)))
		Name := Mgo.GetMsg(mongo,MsgID)
		log := &Mgo.Log{
			Uid:   Uid,
			MsgId: MsgID,
			Name:  Name,
			Param: "",
			Date:  time.Now(),
			Msg:   result.Msg,
		}
		Mgo.InsertLog(mongo,log)
		fmt.Println("有错误消息")
	}else{

	}
	fmt.Println(MsgID)
	//MsgRev.GetData()
	return MsgRev
}