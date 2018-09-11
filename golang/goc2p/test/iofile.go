package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	contents,err := ioutil.ReadFile("a.txdt")
	if err != nil {
		fmt.Println("22")
	}else {
		fmt.Printf("%s\n",contents)
	}
}