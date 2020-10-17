package main

import (
	"fmt"
	"time"
)

/*
上面代码，我们用 for-range 来读取 channel的数据，for-range 是一个很有特色的语句，有以下特点：

- 如果 channel 已经被关闭，它还是会继续执行，直到所有值被取完，然后退出执行。
- 如果通道没有关闭，但是channel没有可读取的数据，它则会阻塞在 range 这句位置，直到被唤醒。
- 如果 channel 是 nil，那么同样符合我们上面说的的原则，读取会被阻塞，也就是会一直阻塞在 range 位置。
*/

func producer(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("大妈做了第", i, "个面包")
		ch <- i
		time.Sleep(time.Second * time.Duration(1))
	}
}

func consumer(ch chan int, count int) {
	for value := range ch {
		fmt.Println("大叔吃了第", value, "个面包")
		count--
		if count == 0 {
			fmt.Println("没面包了， 大叔也吃饱了")
			close(ch)
		}
	}
}

func main() {
	ch := make(chan int)
	count := 5
	go producer(ch, count)
	consumer(ch, count)
}
