package main

import (
	"fmt"
	"io"
	"math"
)

func sqrt(x float64) string{
	if x < 0 {
		return sqrt(-x)+"i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func chengfang( x int,y int) int {
	return x*y
}

func Pic(dx,dy int) [][]uint8 {
	var pic [][]uint8;
	for i:=0;i<dy;i++{
		var one_line []uint8
		for j:=0;j<dx;j++ {
			one_line = append(one_line,(uint8(i & j)))
		}
		pic = append(pic,one_line)
	}
	return pic
}

type rot13Reader struct {
	r io.Reader
}

func (self rot13Reader) rot13Reader(buf []byte) (int,error)  {
	length,err := self.r.Read(buf)
	if err != nil{
		return length,nil
	}
	for i:=0;i<length;i++{
		v :=buf[i]
		switch {
		case 'a' <= v && v<= 'm':
			fallthrough
		case 'A' <=v && v <= 'M':
			buf[i] = v + 13
		}

	}

	return length,nil
}

type adder interface {
	add(string) int
}
type handler func(name string)

func (h handler) add(name string) int{
	return h(name) + 10
}

func process(a adder)  {
	a.add("abc")
}

func dobuler(name string) int  {
	return len(name) * 2
}

type myint int

func (i myint) add(name string) int {
	return len(name)+int(i)
}
func main()  {
	//fmt.Println(sqrt(9))
	//z:= chengfang(6,7)
	//fmt.Println(z)
	var my handler = func(name string) int{
		return len(name)
	}
	fmt.Println(my("difuwei"))
}