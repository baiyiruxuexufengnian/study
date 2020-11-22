package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"sync/atomic"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func main() {
	defer ants.Release()
	runTimes := 10
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < runTimes; i ++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines:%d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
