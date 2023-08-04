package service

import (
	"encoding/json"
	"fmt"
	"github.com/goTouch/TicTok_SimpleVersion/controller"
	"io"
	"net"
	"sync"
)

var chatConnMap = sync.Map{}

func RunMessageServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {
		//一个用户一个conn
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:]) //n :存储读取的字节数;  buf[:] 代表对 buf 的整个切片进行操作
		if n == 0 {
			if err == io.EOF { //is the error returned by Read when no more input is available.
				break //error leads to break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue //not an error ,continue
		}

		var event = controller.MessageSendEvent{} //表示通过调用 controller.MessageSendEvent 的零值构造函数来创建一个新的结构体变量,也可以赋初始值
		_ = json.Unmarshal(buf[:n], &event)       //data from netConnect: buf -> event 网络连接中读取到的数据 buf[:n] 解析为 event 变量所代表的结构体。
		fmt.Printf("Receive Message：%+v\n", event)

		//fmt.Sprintf 将格式化后的字符串作为返回值返回，而 fmt.Printf 则直接将格式化后的字符串输出到标准输出（通常是终端）。
		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			fmt.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		pushEvent := controller.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}
