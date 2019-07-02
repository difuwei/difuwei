package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type Element interface {
	
}
type List []Element

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return "(name:"+p.name+"age)"+strconv.Itoa(p.age)
}

func count(s string)  {
	m := make(map[string]int)
	a := strings.Split(s," ")
	for _,v := range a{
		m[v]+=1
	}
	//fmt.Println(reflect.TypeOf(a))
	fmt.Println(m)
}
type MyFloat float64
func main()  {
	list := make(List,3)
	list[0] = 1
	list[1] = "difuwei"
	list[2] = Person{"zhangsan",18}

	for _,value := range list{
		fmt.Println(reflect.TypeOf(value))
		fmt.Println(reflect.ValueOf(value))
		//fmt.Println()
	}

	count("a a cc a cc di fu wei di di")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)
	//var a = [10]int{}
	//for i:=0 ;i<10; i++{
	//	go func(i int) {
	//		for{
	//			a[i]++
	//			runtime.Gosched()
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
