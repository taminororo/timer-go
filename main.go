package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("カウントダウン開始！")

	ticker := time.NewTicker(1 * time.Second)
	count := 10

	// TODO(human): for range ticker.C { ... } で 10→0 のカウントダウンを書いてください。
	//   - 毎 tick で count を fmt.Println で表示し、count を 1 ずつ減らす
	//   - 0 を表示し終えたら break でループを抜ける
	//   - ループを抜けたら ticker.Stop() で裏方タイマーを止める
	//   ヒント: 「いつ break するか / いつ減らすか」の順番で、0 が出るか出ないかが変わります。
	//           まず動かして、出力を見て調整してみてください。

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
