package main

import (
	"config"
	"fmt"
	"testif"
)

const str1  = "abc"
func main()  {
	config.LoadConfig()
	var a int = 35
	b := string(a)
	fmt.Println(b)
	fmt.Println(str1)
	//testif.TestIf()
	//testif.XunHuan()
 	//fmt.Println(testif.Sqrt(8))
 	//testif.Testreader()
 	//testif.TestImage()
 	go testif.TestGoroutine("world")
	testif.TestGoroutine("hello")
	//c := []int{3,5,2,0,8,4}
	//c1 := make(chan int)
	//go testif.TestChan(c[:len(c)/2],c1)
	//go testif.TestChan(c[len(c)/2:],c1)
	//x,y := <-c1,<-c1
	//fmt.Println(x,y)

}
