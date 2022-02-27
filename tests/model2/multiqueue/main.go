package main
import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var fmeng = false

func produce(id int, wg *sync.WaitGroup, chs chan string){
	cnt := 0
	for !fmeng{
		time.Sleep(1 * time.Second)
		cnt ++
		data := strconv.Itoa(id)+"_"+strconv.Itoa(cnt)
		fmt.Println("produce:",data)
		chs <- data
	}
	wg.Done()
}

func consume(wg *sync.WaitGroup, chs chan string){
	for data := range chs{
		time.Sleep(2 * time.Second)
		fmt.Println("consume",data)
	}
}

func main() {
	//多生产者多消费者
	chans := make(chan string, 10)
	//生产者和消费者waitGroup
	wgP := new(sync.WaitGroup)
	wgC := new(sync.WaitGroup)

	//多生产
	for i:= 0; i<3; i++{
		wgP.Add(1)
		go produce(i,wgP,chans)
	}
	//多消费
	for j:=0;j<2;j++{
		wgC.Add(1)
		go consume(wgC, chans)
	}
	// 制造超时
	go func() {
		time.Sleep(time.Second * 3)
		fmeng = true
	}()

	wgP.Wait()
	close(chans)
	wgC.Wait()

}