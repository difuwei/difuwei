package main

import "fmt"
import "mock"

type Retriver interface {
	Get(url string) string
} 
func downLoad(r Retriver) string{
	return r.Get("www.imooc.com")
}
func main()  {
	var r Retriver
	r = mock.Retriver{"fake imooc"}
	fmt.Println(downLoad(r))
}
