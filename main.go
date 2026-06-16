package main

import (
	"context"
	"fmt"
)

func main() {
	// 手動でキャンセルできる context を作る（cancel() を呼ぶまで生き続ける）
	ctx, cancel := context.WithCancel(context.Background())

	// TODO(human): ctx.Err() の変化を観察する実験を書いてください。
	//   1) cancel() を呼ぶ "前" に ctx.Err() を表示  → 何が出る？
	//   2) cancel() を呼ぶ
	//   3) cancel() を呼んだ "後" に ctx.Err() を表示 → 何に変わった？
	//   ヒント: fmt.Println("キャンセル前 Err():", ctx.Err()) のように出すと変化が見やすい。
	//           生きている context の Err() は nil。
	defer cancel()
	fmt.Println("errは出ない", ctx.Err())
	cancel()
	fmt.Println("errは出ました", ctx.Err())

}
