package main

import "fmt"

func demo() (result int) { // named return value
    result = 10
    defer func() {
        result = result * 2  // تغییر مقدار خروجی قبل از return
    }()
    return // در این لحظه result = 10 اما defer بعد از return اجرا می‌شود
}

func main() {
	x := 10
	defer fmt.Println(x) // مقدار 10 در این لحظه ارزیابی می‌شود
	x = 20
	fmt.Println("x now is:", x)

	fmt.Println(demo())
}