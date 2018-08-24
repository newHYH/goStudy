package main
import "fmt"
func main(){
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	newInt := intSeq()
	fmt.Println(newInt())
	
	p := test(4)
	fmt.Println(p)
	
}
//intSeq函数返回一个匿名函数，匿名函数返回一个int值
func intSeq() func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}
//递归函数练习
func test(n int) int {
	if n == 0{
		return 1
	}
	return n*test(n-1)
}