package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000000; i++ {
		fmt.Println("keep calm and carry on -- chenchen 2")
		time.Sleep(time.Second * 10)
	}
}
