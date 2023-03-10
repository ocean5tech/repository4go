package main

import "fmt"

func main() {
	a := make([]int, 5)

	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	a1 := []int{1, 2, 3, 4, 5, 6}
	printSlice("a1", a1)

	b1 := make([]int, 0, 5)
	printSlice("b1", b1)

	c1 := a1[:2]
	printSlice("c1", c1)

	d1 := c1[2:5] //切片切的都是基础数组，而不是上一个切片
	printSlice("d1", d1)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/* 用 make 创建切片
切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。

make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

a := make([]int, 5)  // len(a)=5
要指定它的容量，需向 make 传入第三个参数：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4 */
