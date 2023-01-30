package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	sqrtv := float64(1)
	for i := 0; i < 10; i++ {
		sqrtv -= (sqrtv*sqrtv - x) / (2 * sqrtv)
		fmt.Printf("sqrtv =  %v  and times of loop = %v\n", sqrtv, i)
	}
	return sqrtv
}

func main() {
	fmt.Println(Sqrt(20))
}

/* 用牛顿法实现平方根函数
计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，
我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：
z -= (z*z - x) / (2*z)
重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。
在提供的 func Sqrt 中实现它。 */
