package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket1 = 10

var wg1 sync.WaitGroup
var matex1 sync.Mutex

func main()  {
	wg1.Add(4)
	go saleTickets1("售票口1")
	go saleTickets1("售票口2")
	go saleTickets1("售票口3")
	go saleTickets1("售票口4")
	wg1.Wait()

}

func saleTickets1(name string)  {
	rand.Seed(time.Now().UnixNano())
	defer wg1.Done()

	for {
		matex1.Lock()
		if ticket1 > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出:", ticket1)
			ticket1 --
		} else {
				matex1.Unlock()
				fmt.Println(name,"售罄，没有票了。。")
				break
		}
		matex1.Unlock()
	}
}
