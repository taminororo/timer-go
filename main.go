package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("カウントダウン開始！")

	ticker := time.NewTicker(1 * time.Second)
	count := 10
	for range ticker.C {

		count -= 1
		fmt.Println(count)
		if count == 0 {
			ticker.Stop()
			break
		}
	}

	fmt.Println("発射！🚀")
}
