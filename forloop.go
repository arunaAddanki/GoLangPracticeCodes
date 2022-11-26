// package main
// import "fmt"

// func main(){
// 	numbers:=[] int{1,2,3,4,5,6,7,8,9,10}
// 	for r:=0;r<10;r++{
// 		fmt.Println(numbers[r])
// 	}
// }


// package main 
// import "fmt"

// func main(){
// 	flowers:=[] string{"lilly","jasmine","sunflower","lotus","marigold"}
// 	for s:=0;s<len(flowers);s++{
// 		fmt.Println(flowers[s])
// 	}
// }


// package main 
// import "fmt"

// func main(){
// 	for t:=0;t<20;t++{
// 		if t==10{
// 			break
// 		}
// 		fmt.Println(t)
// 	}
// }


// package main 
// import "fmt"

// func main(){
// 	for t:=0;t<20;t++{
// 		if t==10{
// 			continue
// 		}
// 		fmt.Println(t)
// 	}
// }

package main 
import "fmt"

func main(){
	for t:=0;t<20;t++{
		if t==10{
			goto label
		}
		fmt.Println(t)
	}
    label:fmt.Println("www.hackerrank.com")
}