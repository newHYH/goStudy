package main
import "fmt"
func main(){
	a := 25
    var b *int
    if b == nil {
        fmt.Println(b)
        b = &a
        fmt.Println(b)
    }
	
	//改变值
	c := 1
	d := &c
	fmt.Println(c)
	fmt.Println(d)
	*d ++
	fmt.Println(c)
	fmt.Println(d)
	//指针参数
	fmt.Println("-----")
	s := 1
	fmt.Println(s)
	change(&s)
	fmt.Println(s)
}
func change(p *int){
	*p = 55;
}