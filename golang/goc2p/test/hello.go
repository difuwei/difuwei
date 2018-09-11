package main

import (
	"fmt"
)

var m = 3
//在函数外不能用这种方式初始化变量
//n := 4

func testVariable()  {

	//变量初始化
	var a int = 222
	b := 3333
	//初始化字典 ①
	var map1 = map[int]string{}
	map1[1] = "mp1"
	map1[2] = "mp2"
	//②
	map2 := map[string]string{"a":"1","b":"2"}
	//③
	map3 := make(map[int]string)
	map3[1] = "aa"
	map3[2] = "bb"
	//④
	map4 := make(map[int]string,5)
	map4[0] = "aaa"
	fmt.Print(a,b,map1,map2,map2["b"])
}

func constTest()  {
	const a,b = 3,4
}


func main() {
	testVariable()
	constTest()
}
