package main

import "fmt"

func main() {
	const pi = 3.14
	const ro = 90

	const (
		he = 10
		ha = 20
	)

	//条件语句
	// a := 100 //int
	// if a == 10 {
	// 	fmt.Printf("b==9")
	// } else if a == 9 {
	// 	fmt.Printf("b=9")
	// } else if a == 8 {
	// 	fmt.Printf("a==8")
	// } else {
	// 	fmt.Printf("a1000")
	// }
	//Switch语句
	// var ss int = 70
	// switch ss {
	// case 90:
	// 	fmt.Printf("优秀")
	// case 80:
	// 	fmt.Printf("良好")
	// case 50, 60, 70:
	// 	fmt.Printf("一般")
	// case 40:
	// 	fmt.Printf("成绩小于50")
	// default:
	// 	fmt.Printf("差")
	// }

	//for循环
	// p := 0
	// for i := 1; i <= 100; i++ {
	// 	p += i
	// 	if i%2 == 0 {
	// 		fmt.Println(i, p)
	// 	}
	// 	if i > 20 {
	// 		break
	// 	}

	// }
	// fmt.Println(p)

	// lable1:
	// 	fmt.Println("这是一个标签")
	// 	p := 0
	// 	for i := 1; i <= 100; i++ {
	// 		p += i
	// 		if i%2 == 0 {
	// 			goto lable1
	// 		}
	// 		// if i > 20 {
	// 		// 	fmt.Println(i, p)
	// 		// }

	// 	}
	// 	fmt.Println(p)
	//数组
	// var blanc = [...]int{1, 2, 3, 4, 5, 6}
	// str := "100000000"
	// fmt.Println(len(blanc), len(str))

	var slice = make([]int, 3, 10)
	slice = append(slice, 2)
	slice = append(slice, 0)
	slice2 := make([]int, len(slice), (cap(slice))*2)
	printSlice(slice)
	copy(slice2, slice)
	printSlice(slice2)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
