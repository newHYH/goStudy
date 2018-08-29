package main
import "fmt"
//结构体
type person struct{
	name string 
	age int
}
func main(){
	fmt.Println(person{"hou",26})
	
	fmt.Println(person{name:"houhou",age:22})
	
	fmt.Println(person{name: "test"})
	
	fmt.Println(&person{name: "hou2", age: 26})
	
	s := person{name: "hou3", age: 50}
    fmt.Println(s.name)
	
	//使用指针才能修改
	sp := &s
    fmt.Println(sp.age)
	
	
	sp.age = 51
    fmt.Println(sp.age)
	
	//这种情况无法修改
	sp2 := s
	sp.age = 52
    fmt.Println(sp2.age)
}