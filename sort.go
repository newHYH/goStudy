package main
import "fmt"
import "sort"
func main(){
	strs := []string{"a","c","b"}
	sort.Strings(strs)
	fmt.Println(strs)
	
	ints := []int{7,2,6}
	sort.Ints(ints)
	fmt.Println(ints)
	
	str2 :=[]string{"a","c","A"}
	sort.Strings(str2)
	fmt.Println(str2)
}