package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Println("swap a", a, "b", b)
}

func swap2(x, y *int) {
	//*x, *y = *y, *x   //这样在main函数中才会交换,main2
	/* swap a 20 b 10
	main1 a 10 b 20
	swap2 x 20 y 10
	main2 a 20 b 10 */
	fmt.Println("swap2 x", *x, "y", *y)
	fmt.Println("bf swap2 x", x, "y", y)
	x, y = y, x //这样在main函数中不会交换,main2
	fmt.Println("af swap2 x", x, "y", y)
	/* swap a 20 b 10
	main1 a 10 b 20
	swap2 x 20 y 10
	main2 a 10 b 20 */
	fmt.Println("swap2 x", *x, "y", *y)
}
func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Println("main1 a", a, "b", b)
	swap2(&a, &b)
	fmt.Println("main2 a", a, "b", b)
}
