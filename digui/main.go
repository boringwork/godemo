package main

import (
	"fmt"
)

func main() {
	a := Feed{RatioMin: 0.1, RatioMax: 0.5}
	b := Feed{RatioMin: 0.1, RatioMax: 0.5}
	c := Feed{RatioMin: 0.1, RatioMax: 0.5}
	d := Feed{RatioMin: 0.2, RatioMax: 0.5}
	e := Feed{RatioMin: 0.2, RatioMax: 0.5}
	f := Feed{RatioMin: 0.2, RatioMax: 0.5}
	feeds := []Feed{a, b, c, d, e, f}
	persent(feeds, 6, 0.01)
}

// Feed df
type Feed struct {
	Ratio    float64
	RatioMax float64
	RatioMin float64
}

// Result sd
type Result struct {
}

func persent(feeds []Feed, count int, step float64) {
	if count == 0 {
		r := 0.0
		for _, f := range feeds {
			r += f.Ratio
		}
		if r-1.0 >= 0.00001 || 1.0-r >= 0.00001 {
			// 总比例超过1，或不够1，不符合
		} else {
			// 满足比例，进行营养计算
			calc(feeds)
			// 然后与标准值对比
			for i, f := range feeds {
				fmt.Printf("[%d]=%2.2f, ", i, f.Ratio)
			}
			fmt.Println()

		}

		return
	}

	index := len(feeds) - count
	for feeds[index].Ratio = feeds[index].RatioMin; feeds[index].Ratio <= feeds[index].RatioMax; feeds[index].Ratio += step {

		persent(feeds, count-1, step)
	}
}

func calc(feeds []Feed) Result {
	// TODO 利用公式计算出结果，然后与标准值对比
	return Result{}
}
