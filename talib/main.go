package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

// 株価分析
func main() {
	// yahooの株価を取得
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-04-01", "2019-01-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	mva := talib.Ema(spy.Close, 14) // 14日間
	fmt.Println(mva)
}
