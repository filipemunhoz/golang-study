package main

import (
	"fmt"
	"time"
)

func mainjk() {

	for timer := 10; timer >= 0; timer-- {

		if timer == 0 {
			fmt.Println("Timesup!")
			break
		}
		fmt.Println("time:", timer)
		time.Sleep(1 * time.Second)
	}

}
