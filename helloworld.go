package main

import (
	"fmt"
	"sync"
)
// 设置 schedtrade=X 会使调度器每X毫秒打印一行标准错误，汇总调度器的状态
func main() {
	c := make(chan int, 10)
    fmt.Println(c)
	wg := sync.WaitGroup{}
   wg.Add(10)
   for i := 0; i < 10; i++ {
       go func(wg *sync.WaitGroup) {
           var counter int
           for i := 0; i < 1e10; i++ {
               counter++
           }
           wg.Done()
       }(&wg)
   }

   wg.Wait()
}