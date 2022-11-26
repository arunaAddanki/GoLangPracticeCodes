package main 
import "fmt"

func main(){
	Dresses:=[] string{"skirts","jeans","lehenga","tops","sarees","blazors","jerkins"}
	for q:=range Dresses{
		fmt.Println(Dresses[q])
	}
}