package Result

import (
	"github.com/golang/protobuf/proto/test_proto"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main() {
	
}
//创建公司结果
func CreateCompanyResult(Data []byte) *AutoMsg.CreateCompanyResult{
	CreateCompanyResult := &AutoMsg.CreateCompanyResult{}
	err := proto.Unmarshal(Data, CreateCompanyResult)
	if err!= nil {
		fmt.Println("解析失败CreateCompanyResult")
	}
	return CreateCompanyResult
}
