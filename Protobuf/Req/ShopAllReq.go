package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
)

//商店请求
func ShopAllReq(MsgId int32) []byte {
	Shop := &AutoMsg.ShopAllReq{}
	data, err := proto.Marshal(Shop)
	if err != nil {
		fmt.Println("1")
	}
	return Send(MsgId, data)

}
