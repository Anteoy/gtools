// 如果文件夹也是main
//# _/home/zhoudazhuang/class100/gtools/test/base/note/gotest/main
//runtime.main_main·f: relocation target main.main not defined
//runtime.main_main·f: undefined: "main.main"
//并且文件夹必须为*_test.go
//不正确的goroot会引起下面fail: GOPATH这里倒是可以不用设置,go test知道测试当前这个文件夹(或者说包)
///tmp/go-build066857360/_/home/zhoudazhuang/class100/gtools/test/base/note/gotest/notmain/_test/_testmain.go:54:24: cannot use matchString (type func(string, string) (bool, error)) as type testing.testDeps in argument to testing.MainStart:
//func(string, string) (bool, error) does not implement testing.testDeps (missing ImportPath method)
//$ echo $GOROOT
///home/zhoudazhuang/usr/local/go1.9.4
//??? 注意: 本地环境变量的GOROOT居然和go version不是一个版本 导致的这个问题 就像我的网卡有时候不能正确识别一样 经过测试发现是在gogland下的命令行会出现的这个问题 在terminal下不会出现这个路径不一致的问题.
//# zhoudazhuang at zhoudazhuang-pc in ~/class100/gtools/test/base/note/gotest/src/notmain on git:master x [10:27:40]
//$ go version
//go version go1.7rc5 linux/amd64

// test oper:
// go test -v
// go test -v -timeout 3s
//go test -v -run "TestSum"
//go test -v -run "(?i)sum"
// go test -v -bench .
//-benchmem 输出内存统计信息。
// -benchtime t 设置每个性能测试运⾏时间。 默认是测1秒执行的次数和平均执行一次的时间
// -cpu 设置并发测试。默认 GOMAXPROCS。
//go test -bench . -benchmem -cpu 1,2,4 -benchtime 30s (经测试-run无效)
//go test -v 同时会测试example

// go test -cover -coverprofile=cover.out -covermode=count
// go test -cover
package notmain

import (
	"fmt"
	"testing"
	"time"
)

func sum(n ...int) int {
	var c int
	for _, i := range n {
		c += i
	}
	return c
}
func TestSum(t *testing.T) {
	time.Sleep(time.Second * 2)
	if sum(1, 2, 3) != 6 {
		t.Fatal("sum error!")
	}
}
func TestTimeout(t *testing.T) {
	time.Sleep(time.Second * 5)
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if sum(1, 2, 3) != 6 {
			b.Fatal("sum")
		}
	}
}

func ExampleSum() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20, 30))
	// Output:
	// 6
	// 60
}
