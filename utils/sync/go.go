package sync

import "fmt"

func Go(x func()) {
	if err := recover(); err != nil {
		fmt.Println("go 携程裂开了")
		//  todo 报警消息通知,有必要记录日志
	}
	go x()
}
