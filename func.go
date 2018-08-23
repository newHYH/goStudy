package main
import "fmt"
func main(){
	res := add(1,2)
	fmt.Println(res)
	
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_,c := vals()
	fmt.Println(c)
	
	sum(1,2,3)//正常的参数可以传递
	nums := []int{1,2,3,4}
	sum(nums...)//参数是切片时，实参传递要用...
	
	
}
//函数返回一个值
func add(a int,b int)int{
	return a+b
}
//函数返回多个值
func vals()(int,int){
	return 1,2
}
//可变参数：使用切片slice...作为参数
func sum(nums...int){
	fmt.Println(nums)
	sum := 0
	for _,num :=range nums{
		sum += num
	}
	fmt.Println(sum)
}