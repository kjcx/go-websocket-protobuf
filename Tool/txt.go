package Tool

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
	"io"
	"WebSocket/Mgo"
)
//解析txt并发插入mongo数据库
func Txt() {
	fileName := "./a.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		arr := strings.Split(line," ")
		fmt.Println(arr[0],arr[1],arr[2])
		//插入数据库
		mongo := Mgo.Mongo()
		msgid,err :=strconv.Atoi(arr[0])
		Mgo.MsgProto(mongo,&Mgo.MsgInfo{int32(msgid),arr[1],arr[2]})

		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}

