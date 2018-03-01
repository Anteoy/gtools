package main

import(
	"time"
	"fmt"
	"math/rand"
)

func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		// sub  jian ge shi jian
		fmt.Println(time.Now().Sub(start))
		// she zhi zai randomDuration() zhi hou guo qi
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}