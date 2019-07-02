package main

import "fmt"

func fibonaci() func() int {
	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}

func trydefer()  {
	for i:=0;i<10;i++ {
		defer fmt.Println(i)
	}
}
func main()  {
	trydefer()
	//fmt.Println(fibonaci()())
}