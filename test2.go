package main
import "fmt"
func main(){
	//go语言切片练习
	s:= make([]string,3)
	fmt.Println(s)
	fmt.Println(len(s))
	//set方法
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	//append方法
	s = append(s,"d")
	fmt.Println(s)
	c := make([]string,len(s))
	fmt.Println(c)
	//复制切片
	copy(c,s)
	fmt.Println(c)
	l := s[2:5]
	k := s[2:4]
	m := s[2:3]
	fmt.Println(l)
	fmt.Println(k)
	fmt.Println(m)
	///////////////////////
	//数组求和range用法
	nums := []int{1,2,3}
	sum := 0
	for _,num := range nums{
		sum += num
	}
	fmt.Println(sum)
	
}
























