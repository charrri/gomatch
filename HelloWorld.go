package main
import "fmt"

func main_wt_1() {
	fmt.Println("Hello, World!")
	
	var a int32 = 1
	var b = a
	fmt.Println(a,b)

	b = 2
	fmt.Println(a,b)

	Swap(a,b)
	fmt.Println(a,b)

	a, b = Swap(a,b)
	fmt.Println(a,b)
}

func Swap(a,b int32) (int32, int32) {
	return b, a
}
