package main
import "fmt"
func main() {
	fmt.Println("Enter size of your array: ")        //taking user input from array
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)           
	for i := 0; i < size; i++ {
		// fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is: ", arr)
}