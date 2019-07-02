package testif

import (
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

func TestIf()  {
	if a := 3;a > 2{
		fmt.Println(a)
	}
}
func XunHuan()  {
	a :=1
	for a < 5{
		fmt.Println(a)
		a++
	}
}

func Sqrt(x float64) (float64,error) {
	if x > 0 {
		return math.Sqrt(x),nil
	}

	return x,nil
}

func Testreader()  {
	r := strings.NewReader("hell,world")
	b := make([]byte,2)
	for{
		n,err := r.Read(b)
		//fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func TestImage(){
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
}

func TestGoroutine(s string){
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func TestChan(a []int,c chan int)   {
	sum := 0

	for _,v := range a{
		sum += v
	}
	c <- sum
}

func TestFibonaci(n int,c chan int){
	x,y := 0,1
	for i:=0; i<n; i++ {
		c <- x
		x,y = y,x+y
	}
	close(c)
}