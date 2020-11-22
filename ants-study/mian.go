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
	fmt.Printf("runing with : %d\n", n)
}

func main() {
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	for i := 1; i <= 10; i ++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("runing goroutine counts : %d\n", p.Running())
	fmt.Printf("finish all task, result:%v\n", sum)
}
