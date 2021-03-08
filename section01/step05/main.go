// STEP05: レア度を作ってみよう

package main

import (
	"math/rand"
	"time"
)

func main() {

	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	// 0から3までの間で乱数を生成する
	num := rand.Intn(4)

	// TODO: 変数numが0のときは"ノーマル"、
	// 1のときは"R"、2のときは"SR"、
	// 3のときは"XR"と表示する
	if num == 0 {
		println("ノーマル")
	} else if num == 1 {
		println("R")
	} else if num == 2 {
		println("SR")
	} else if num == 3 {
		println("XR")
	}

}
