package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"log"
)

//解析登录返回 1027
func ConnectingResult(Data []byte) *AutoMsg.ConnectingResult {
	Connecting := &AutoMsg.ConnectingResult{}
	err := proto.Unmarshal(Data, Connecting)
	if err != nil {
		log.Fatal("marshaling error: ")
	}
	return Connecting
}
