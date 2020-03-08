package algorithm

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	wg   = sync.WaitGroup{}
	mu   = &sync.Mutex{}
	cond = sync.NewCond(mu)
)
// 启动两个线程, 一个输出 1,3,5,7…99, 另一个输出 2,4,6,8…100 最后 STDOUT 中按序输出 1,2,3,4,5…100
func TestCond(t *testing.T) {
	wg.Add(2)
	go t2()
	time.Sleep(1 * time.Second)
	go t1()
	wg.Wait()
	fmt.Println("end")
}

func t1() {
	for i := 1; i <= 100; i += 2 {
		mu.Lock()
		fmt.Println("t1:", i)
		cond.Broadcast()
		cond.Wait()
		mu.Unlock()
	}
	wg.Done()
}

func t2() {
	for i := 2; i <= 100; i += 2 {
		mu.Lock()
		cond.Wait()
		fmt.Println("t2:", i)
		cond.Broadcast()
		mu.Unlock()
	}
	wg.Done()
}
