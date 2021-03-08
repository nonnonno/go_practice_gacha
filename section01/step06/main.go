// STEP06: レア度ごとに出る確率を変えてみよう

package main

import (
	"math/rand"
	"time"
)

func main() {

	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())

	// 0から99までの間で乱数を生成する
	num := rand.Intn(100)

	// TODO: 変数numが0〜79のときは"ノーマル"、
	// 80〜94のときは"R"、95〜98のときは"SR"、
	// それ以外のときは"XR"と表示する
	if num >= 0 && num <= 79 {
		println("ノーマル")
	} else if num >= 80 && num <= 94 {
		println("R")
	} else if num >= 95 && num <= 98 {
		println("SR")
	} else {
		println("XR")
	}
}
