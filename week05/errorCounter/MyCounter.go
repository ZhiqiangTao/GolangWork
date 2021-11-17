package week05

import (
	"fmt"
	"golangwork/week05"
	"time"
)

type MyCounter struct {
	//假设我的计数器是统计错误次数的
	Errors *week05.Counter
}

func (c *MyCounter) Update(newVal float64) {
	c.Errors.Increment(newVal)

	fmt.Println(time.Now(), "此次update结束，error计数器结果如下")
	for key, value := range c.Errors.Buckets {
		fmt.Printf("			%v, %v\n", key, value.Value)
	}
}
