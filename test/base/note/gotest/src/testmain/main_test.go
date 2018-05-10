package testmain

import (
	"log"
	"os"
	"testing"
)

//用于测试的Main函数 go test -v
//go test -v -run TestA 同样可以限制指定的函数
//go test -v -bench .
func TestMain(m *testing.M) {
	println("setup")
	code := m.Run()
	println("teardown")
	os.Exit(code)
}
func TestA(t *testing.T) {
	log.Println("[TestA] running ")
}

//功能测试
func TestB(t *testing.T) {
	log.Println("[TestB] running ")
}

//基准测试  Go的Benchmark默认是在1s的时间内,尽可能多的执行这个函数. 可以通过 -benchtime 这个参数扩大时长
func BenchmarkC(b *testing.B) {
	log.Println("[BenchmarkC] running ")
}

//样本测试
func ExampleC() {
	log.Println("[ExampleC] running ")
}
