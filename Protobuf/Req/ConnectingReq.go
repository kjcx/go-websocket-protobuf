package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
	"log"
	"WebSocket/Protobuf"
)

//发起连接

func ConnectingReq(MsgID int32, Token string) []byte {
	test := &AutoMsg.ConnectingReq{
		//Token: "token:5405ef6818c54e0ec07430fb1ad968ba",
		Token: Token,
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)
	return SendRev.Send(MsgID, data)

}
