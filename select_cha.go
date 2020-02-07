package main 

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)
	quit := make(chan bool)

	//goroutine
	go func(){
		for{
			select{
				case num := <-ch:
					fmt.Println("nums=",num)
				case <-time.After(3 * time.Second):
					fmt.Println("超时")
					quit <- true  //写入
			}
		}
	}()

	for i := 0;i<5;i++{
		ch <-i
		time.Sleep(time.Second)
	}

	<-quit //这里暂时阻塞，直到可读
	fmt.Println("程序结束")
}
