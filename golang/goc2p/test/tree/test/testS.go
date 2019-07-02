package main

import "fmt"
import "tree"

type person struct {
	firstName string
	lastName string
	age     int
}

func (this *person) addPerson(data string)  {
	this.firstName = "wang"
	this.lastName = "xiaoming"
	this.age = 20
}

//不定个数参数
func add (numbers ...int) int{
	sum := 0
	for _,j:= range numbers{
		sum +=j
	}

	return sum
}

func main()  {
	fmt.Println(tree.ConstTest())
}