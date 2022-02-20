package model1

import (
	"fmt"
	"testing"
	"time"
)
//模块1

/**
	作业1.1

	编写一个小程序：
	给定一个字符串数组
	[“I”,“am”,“stupid”,“and”,“weak”]
	用 for 循环遍历该数组并修改为
	[“I”,“am”,“smart”,“and”,“strong”]

	测试：go test -v tests/model1/model1_test.go  -run TestArray
 */
func TestArray(t *testing.T) {
	var strArr = []string{"I","am","stupid","and","weak"}
	var newStrArr = []string{"I","am","smart","and","strong"}
	for i:= 0; i<len(newStrArr) ; i++ {
		strArr[i] = newStrArr[i]
	}

	fmt.Println("new array is",newStrArr)
}

/**
	作业1.2

	基于 Channel 编写一个简单的单线程生产者消费者模型：

	队列：
	队列长度 10，队列元素类型为 int
	生产者：
	每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
	消费者：
	每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

	测试：
 */
func TestQueueChan(t *testing.T) {
	len := 10
	var ch = make(chan int, len)
	defer func(ch chan int) {
		close(ch)
		fmt.Println("channel is closed")
	}(ch)

	producer := func(ch chan int) {
		for i:=0; i<len; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}

	//异步生产
	go producer(ch)
	//同步阻塞消费
	num :=0
	for {
		if (num>=len){
			break;
		}

		select {
			case data:= <-ch:
				num++
				fmt.Println(data);
			default:
				fmt.Println("wait")
		}
		time.Sleep(1 * time.Second)
	}

}