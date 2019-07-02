package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
)

func main() {
	for i :=0; i<10;i++  {

	}
	//struct
	//type Person struct {
	//	id int
	//	name string
	//	Age int
	//	Address string
	//}
	//person类型为结构体Person类型
	//person := Person{id:11,name:"trump"}
	//person1是指针变量
	//person1 := new(Person)

	//arr1 := []int{3, 4, 3, 5, 6, 8, 10, 23}
	//fmt.Println(cap(arr1))

	//for index := 0; index < len(arr1); index++ {
	//	fmt.Println(index)
	//}
	//for _,value := range arr1{
	//	//fmt.Println(value)
	//}
	//arr1 = append(arr1, 44444)
	//str1 := "aabbedddd"
	//_, ok := interface{}(str1).(string)
	//fmt.Println(ok)

	//a := 106
	//f := func() int{ a= a*2;return 5}
	//fmt.Println(a)
	//x := []int{a,f()}
	//一个汉字3个字节
	//str2 := "猴子\\看"
	//fmt.Println(str2)
	//testArray()
	//testArr2()
	//testSlice()

	//testMap()
	fmt.Println(testFunc(38))
	fmt.Println(testFunc(60))
	testfunc := func(b int){
		fmt.Println(b)
	}
	testfunc(20)

	done := make(chan struct{})
	langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
	for _, l := range langs {
		go hello(l, done)
	}

	for _ = range langs {
		<-done
	}

	p := Person1{
		20,
		attr{
			name:"tom",
			age:20,
		},
	}
	p.name = "lucy"
	fmt.Println(p)
	fmt.Println(os.Getpid())

	testPipe()

	ppp := Person1{}
	fmt.Println(ppp.name)
}

func testArray()  {
	//数组长度是类型的一部分 [5]int{}和 [4]int{}不是同一种类型
	arr1 := [5]int{}
	fmt.Println(arr1)

}
func testArr2()  {
	arr1 := [5]string{"bac","aaa","bbb","ccc"}
	var arr2 = [5]string{}
	for k,v := range arr1{
		arr2[k] = v
	}

	arr3 := [10]int{1,2,3,4,52,3,1,5,8}
	slice3 := arr3[2:5]
	fmt.Println(cap(slice3))
	fmt.Println("aaaaaaaaa")
}
func testSlice()  {
	slice1 := []string{"name","id","age"}
	slice1 = append(slice1,"java")
	fmt.Println(slice1)
	s := make([]string,10)
	fmt.Println(&s[0])
	fmt.Println(&s[1])
	fmt.Println(&s[2])
	fmt.Println(&s[3])
}
func testMap()  {
	map1 := make(map[int]string)
	map1[2]="abc"
	map1[1]="cddd"
	value1,ok := map1[1]
	fmt.Println(ok)
	if ok {
		fmt.Println(value1)
	}
	fmt.Println(len(map1))
}

func testFunc(a int) bool {
	if a >= 60 {
		return true
	}
	return false
}

type Interface interface {
	Len() int
	Less(i,j int) bool
	Swap(i,j int)
}

//sort.Interface 在一个接口声明中嵌入另一个接口声明：将一个接口中方法批量添加到另一个接口中
type Sortable interface {
	sort.Interface
	Sort()
}

type SortableStrings [3]string

func (self SortableStrings) Len() int  {
	return len(self)
}
func (self SortableStrings) Less(i,j int) bool{
	return self[i] > self[j]
}
func (self SortableStrings) Swap(i,j int)  {
	self[i],self[j] = self[j],self[i]
}

func (self SortableStrings) Sort(){
	sort.Sort(self)
}

type Person struct {
	id int
	name string
}

func hello(name string, done chan struct{}) {
	fmt.Println("hello ",name)
	done <- struct{}{}
}

type attr struct {
	name string
	age int
}
type Person1 struct {
	id int
	attr
}

func testPipe(){
	reader,writer,err := os.Pipe()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(reader,writer,"difuwei")
}

func testConst()  {
	const a  = "这是个常量"
	fmt.Println(a)
}
func testCmd()  {
	cmd0 := exec.Command("echo","-n","my first command")
	//创建一个能够获取此命令的输出管道
	//stdout0 的类型是io.ReadCloser
	stdout0,err := cmd0.StdoutPipe()
	if err != nil{
		fmt.Println(err)
		return
	}
	output0 := make([]byte,30)
	n,err := stdout0.Read(output0)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n",output0[:n])
	if err := cmd0.Start();err != nil{
		fmt.Println(err)
		return
	}
	for n:=1;n>0;n /= 2{
		fmt.Println(n)
	}
}