package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:"comp"` //在json中的别名
	Subjects []string
	IsOk bool `json:",string"` //在json中 将bool类型转换为string类型
	Price float64 `json:"-"` //在json中不输出该字段
}

type A struct {
	RoleIds []string
}
func main(){
	str := `{"IsOk":true,"Subjects":["A","B","C"],"company":"KingSoft"}`
	m := make(map[string]interface{})
	json.Unmarshal([]byte(str),&m)
	fmt.Println(m)

	i := &A{} //new一个结构体出来，返回的是指针类型
	str1 := `{"RoleIds":["A","B","C"]}`
	json.Unmarshal([]byte(str1),i)
	fmt.Printf("%+v",i)
	//输出是 &{Company:KingSoft Subjects:[A B C] IsOk:true Price:0}


}
