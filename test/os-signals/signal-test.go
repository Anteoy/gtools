//Go中通过os/signals包，可以接受系统信号。
package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("awaiting signal")
	<-sigs
	fmt.Println("exiting")
}

// 命令行执行
// 输出结果：
// awaiting signal
// ^Cexiting
// 运行后输awaiting signal，然后卡住。当在键盘上按control+c以后go收到SIGINT信息，向sigs channel中添加信号，<-sigs处停止阻塞，程序执行结束。
// 此特性一般用来正常的退出程序，收到信息号，程序执行一系列的清理工作，然后退出。
