package main

import "fmt"

func main()  {

	// 1、
	m1 := map[string]string{
		"name":"xi",
		"id":"001",
		"course":"math",

	}
	//2、
	m2 := make(map[string]int)

	//3、
	var m3 map[string]int

	fmt.Println(m1,m2,m3)

	for k,v := range m1{
		fmt.Println(k,v)
	}
	fmt.Println(m1["course"])
	delete(m1,"id")

	fmt.Println(m1)
	//寻找最长不含有重复字符的子串
	//abcabcbb->abc




}
