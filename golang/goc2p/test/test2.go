package main

import (
	"fmt"
)
//
//var ch chan  int = make(chan int,5)
//func foo(id int){
//	ch <- id
//}
type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)",p.name,p.age)
}
func main()  {

	a := Person{"difuwe",18}
	z := Person{"xijinping",20}
	fmt.Println(a,z)
	//for i:=0;i<5;i++{
	//	go foo(i)
	//}

	//for i:=0;i<5;i++{
	//	fmt.Println(<-ch)
	//}

	ch1 := make(chan int,5)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	for v:= range ch1 {
		fmt.Println(v)
		if len(ch1) == 0 {
			break
		}
	}
	}

