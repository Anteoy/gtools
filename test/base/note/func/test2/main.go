package main

func add(x, y int) (z int) {
	defer func() {
		z = z + 100
		println(z) // 输出: 303
	}()
	z = x + y
	return z + 200 // 执⾏顺序: (z = z + 200) -> (call defer) -> (ret)
}

// 显式 return 返回前，会先修改命名返回参数。 defer 里面可以修改返回值z
func main() {
	println(add(1, 2)) // 输出: 303
}
