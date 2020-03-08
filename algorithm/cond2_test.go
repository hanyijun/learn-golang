package algorithm

import (
	"fmt"
	"sync"
	"testing"
)

var (
	ch  = make(chan int)
	wg2 = &sync.WaitGroup{}
)

// 启动两个线程, 一个输出 1,3,5,7…99, 另一个输出 2,4,6,8…100 最后 STDOUT 中按序输出 1,2,3,4,5…100
func TestCond2(t *testing.T) {
	wg2.Add(2)
	go t3()
	go t4()
	wg2.Wait()
	fmt.Println("end")
}

func t3() {
	for i := 1; i <= 100; i += 2 {
		if i != 1 {
			<-ch
		}
		fmt.Println("t3:", i)
		ch <- 1
	}
	wg2.Done()
}

func t4() {
	for i := 2; i <= 100; i += 2 {
		<-ch
		fmt.Println("t4:", i)
		if i != 100 {
			ch <- 1
		}
	}
	wg2.Done()
}
