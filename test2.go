package main
import "fmt"
func main(){
	s:= make([]string,3)
	fmt.Println(s)
	fmt.Println(len(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	s = append(s,"d")
	fmt.Println(s)
	c := make([]string,len(s))
	fmt.Println(c)
	copy(c,s)
	fmt.Println(c)
	l := s[2:5]
	k := s[2:4]
	m := s[2:3]
	fmt.Println(l)
	fmt.Println(k)
	fmt.Println(m)
	///////////////////////
	nums := []int{1,2,3}
	sum := 0
	for _,num := range nums{
		sum += num
	}
	fmt.Println(sum)
	
}