
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)
type Stringer interface {
	String() string
}

type Binary uint64

func (i Binary) Stringer() string {
	return strconv.FormatUint(i.Get(),2)

}
func (i Binary) Get() uint64 {
	return uint64(i)
}


func main()  {
	if contents,err := ioutil.ReadFile("a.txt");err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}

	//if err != nil {
	//	fmt.Println("22")
	//}else {
	//	fmt.Printf("%s\n",contents)
	//}
	//
	//t := []int{1,2,3,4}
	//s := make([]interface{},len(t))
	//for i,v := range t{
	//	s[i] = v+1
	//}
	//
	//fmt.Println(s)
	//var b Binary = 32
	////bb := Stringer(b)
	//fmt.Println(b.Stringer())

}
