package Req

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	
}
//创建公司请求
func CreateCompanyReq(Name string) []byte{
	CreateCompanyReq := &AutoMsg.CreateCompanyReq{
		Name:Name,
	}
	Data,_ := proto.Marshal(CreateCompanyReq)
	return Send(1006,Data)
}
