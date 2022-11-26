package main
import "fmt"
func main(){
	var a int = 10000

	var b *int 
	
	b = &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(&a)

}