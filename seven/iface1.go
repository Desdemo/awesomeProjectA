package main

import (
	"fmt"
	"time"
)

func main()  {
	var x interface{} = time.Now()
	fmt.Println(x)
	fmt.Println("%T",x)
}
