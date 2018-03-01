//场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100

package main

import (
	"fmt"
	"time"
	"sync"
)

type Ban struct {
	visitIPs map[string]time.Time
	ml *sync.Mutex
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
var ml sync.Mutex
func (o *Ban) visit(ip string) bool {
	ml.Lock()
	//must first
	defer ml.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	//fatal error: concurrent map read and map write
	//ml.Lock()
	o.visitIPs[ip] = time.Now()

	go delete3(o.visitIPs,ip)
	return false
}

func delete3(visitIPs map[string]time.Time,ip string) {
	time.Sleep(time.Minute * 3)
	//fatal error: concurrent map read and map write
	ml.Lock()
	defer ml.Unlock()
	delete(visitIPs,ip)
}

//fatal error: concurrent map writes
func main() {
	var success int32 = 0
	ban := NewBan()
	var wg  sync.WaitGroup
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				// 3分钟内 此ip是否访问
				if !ban.visit(ip) {
					//atomic.AddInt32(&success,1)
					//no atomic success: 56 dan shi duo shu shi 100
					success++
				}
			}(j)
		}

	}
	wg.Wait()
	fmt.Println("success:", success)
}
