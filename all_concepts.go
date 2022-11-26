// package main

// import(
// 	"fmt"
// "math"
// )

 

// func main() {
// 	var a float64 = 9
 
// 	var result = math.Sqrt(a)
// 	fmt.Println(result)
// }


// package main
// import "fmt"
// func main(){
// 	var fruits=[3]string{"apple","banana","grapes"} //array declaration with length
// 	fmt.Println(fruits)
// }



// package main
// import "fmt"
// func main(){
// 	var icecreams = []string{"vennela","strawberry","butterscotch","anjan"} //array declaration without length
// 	fmt.Println(icecreams)
// }



// package main
// import "fmt"
// func main(){
// 	var dishes = []string{"chicken","mutton","prawns","fish","mushrooms"}
// 	dishes[4] = "crabs"
// 	dishes = append(dishes, "pig","duck","beef")
// 	fmt.Println(dishes)
// }



// package main
// import "fmt"
// func main(){
// 	var breakfast = []string{"idli","dosa","vada","pongal","poha","upma"}
// 	fmt.Println(breakfast[1:4+1])
// 	fmt.Println(len(breakfast))
// 	fmt.Println(cap(breakfast))

// 	breakfast=append(breakfast,"bonda","tomato bath","bisibele bath","lemon rice","carrot rice","tamarind rice","paasta","pani puri")
// 	breakfast = append(breakfast,"puttu","uppam")
// 	fmt.Println(len(breakfast))
// 	fmt.Println(cap(breakfast))

// }




// package main
// import "fmt"
// func main(){
// 	var cakes = []string{"cream","choclate","ice","redvelvet","blackforest","pineapple"}
// 	fmt.Println(cakes[1:4])
// 	var a = (cakes[1:4])
// 	fmt.Println(a)
	
// }



// package main
// import "fmt"
// func main(){
// 	Dresses:=[]string {}
// 	fmt.Printf("%d/n",len(Dresses))
// 	fmt.Printf("%d/n",cap(Dresses))
// 	Dresses = append(Dresses, "palazoo","jeans","jerkins","coats","jeans")
// 	fmt.Println(Dresses)
// 	Dresses[2] = "Skirts"
// 	Dresses[3] = "Shorts"
// 	fmt.Println(Dresses)
// }






// package main
// import "fmt"
// func main(){
// 	var alphabets=[]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o"}
// 	fmt.Println(alphabets)
// 	fmt.Println(len(alphabets))           
// 	fmt.Println(cap(alphabets))

// 	newAlphabets := alphabets[:len(alphabets)-9]
// 	alphabetscopy := make([]string,len(newAlphabets))
// 	copy(alphabetscopy,newAlphabets)
// 	fmt.Println(alphabetscopy)

// }




// package main
// import "fmt"
// func main(){
//    b := 10
// //    var b int=10
// //    var b =10
// 	if b==10{
// 		fmt.Println("true")
// 	}
// }



// package main
// import "fmt"
// func main(){
// 	var a =10
// 	var b = 20
// 	if a>b{
// 		 fmt.Println("true")
// 	}else{
// 		 fmt.Println("false")
// 	}

// }


// package main
// import "fmt"
// func main(){
// 	a:=2000
// 	b:=3000
// 	if a>b{
// 		fmt.Println(a+b)
// 	}else if a<b{
// 		fmt.Println(a-b)
// 	}else{
// 		fmt.Println("exit")
// 	}
// }


// package main
// import "fmt"
// func anjan(){
// 	fmt.Println("anjan")

// }
// func main(){
// 	anjan()
// 	anjan()
// }


// package main
// import "fmt"
// func surname(sname string,age int){
// 	fmt.Println("Hello",sname,"How are you",age) //passing arguments to function

// }

// func main(){
// 	surname("aru",23)
// 	surname("anjan",24)
// 	surname("sriram",24)
// 	surname("naga",23)

// }




// package main
// import "fmt"
// func add(a int,b int) int {
// 	return a+b
	

// }

// func sub(a int ,b int) int {
// 	return a-b
// }
// func main(){
// 	fmt.Println(add(1000,2000),sub(40,30))

// }


// package main
// import ("fmt")

// func testcount(x int) int {
//   if x == 20 {
//     return 0
//   }
//   fmt.Println(x)
//   return testcount(x + 1)
// }

// func main(){
//   testcount(2)
// }




// package main
// import "fmt"
// type human struct{
// 	love bool
// 	affairs bool 
// 	vir bool 
// 	age int 
// }
// func main(){
// 	var anjan human
// 	var aruna human 
	

// 	anjan.love = false
// 	anjan.affairs = false
// 	anjan.vir = false
// 	anjan.age = 24

// 	aruna.love = false 
// 	aruna.affairs = false
// 	aruna.vir = true
// 	aruna.age = 23


// 	fmt.Println(anjan.love)
// 	fmt.Println(anjan.affairs)
// 	fmt.Println(anjan.vir)
// 	fmt.Println(anjan.age)


// 	fmt.Println(aruna.love)
// 	fmt.Println(aruna.affairs)
// 	fmt.Println(aruna.vir)
// 	fmt.Println(aruna.age)
// }


// package main
// import "fmt"

// type Person struct {
//   name string
//   age int
//   job string
//   salary int
// }
//                               //pass struct as functions
// func main() {
//   var pers1 Person
//   var pers2 Person

//   // Pers1 specification
//   pers1.name = "Hege"
//   pers1.age = 45
//   pers1.job = "Teacher"
//   pers1.salary = 6000

//   // Pers2 specification
//   pers2.name = "Cecilie"
//   pers2.age = 24
//   pers2.job = "Marketing"
//   pers2.salary = 4500

//   // Print Pers1 info by calling a function
//   printPerson(pers1)
  
//   // Print Pers2 info by calling a function  
//   printPerson(pers2)
// }

// func printPerson(pers Person) {
//   fmt.Println("Name: ", pers.name)
//   fmt.Println("Age: ", pers.age)
//   fmt.Println("Job: ", pers.job)
//   fmt.Println("Salary: ", pers.salary)
// }





// package main
// import "fmt"                  //map 
// func main(){
// 	var a = map[string]string{"bike":"passion pro","car":"renault","bicycle":"ladybird"}
// 	b:=map[string]int{"aru":23,"anjan":24,"anu":20}
// 	fmt.Printf("%v\n",a)
// 	fmt.Printf("%v\n",b)
// }




// package main
// import "fmt"
// func main(){
// 	var a = make(map[string]string)
// 	 b := make(map[string]string)

// 	a["brand"]="Hero"
// 	a["model"]="passion pro"
// 	a["year"] = "2014"

// 	b["brand"] = "car"
// 	b["model"] ="renault"
// 	b["year"] = "2020"

// 	fmt.Printf("%v\n",a)
// 	fmt.Printf("%v\n",b)

// }





// package main
// import ("fmt")

// func main() {
//   var a = make(map[string]string)
//   var b map[string]string   //creating empty array

//   fmt.Println(a == nil)
//   fmt.Println(b == nil)
// }





// package main
// import "fmt"
// func main(){
// //var a = map[string]string{"bike":"Hero","model":"shine"}
// //b:=map[string]int{"anjan":24,"aruna":23}

// var a = make(map[string]string)
// b:=make(map[string]int)
// a["bike"]="Hero"
// a["model"]="shine"

// b["anjan"]=24
// b["aruna"]=23

// fmt.Printf("%v\n",a)
// fmt.Printf("%v\n",b)
// }






// package main
// import ("fmt")

// func main() {
//   var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

//   val1, ok1 := a["brand"] // Checking for existing key and its value
//   val2, ok2 := a["color"] // Checking for non-existing key and its value
//   val3, ok3 := a["day"]   // Checking for existing key and its value
//   _, ok4 := a["model"]    // Only checking for existing key and not its value

//   fmt.Println(val1, ok1)
//   fmt.Println(val2, ok2)
//   fmt.Println(val3, ok3)
//   fmt.Println(ok4)
// }





// package main 
// import "fmt"
// func main(){
// 	var a int=1000
// 	var p *int   //declaring the pointer
// 	p = &a      // initialising the pointer
// 	fmt.Println(a)
// 	fmt.Println(&a)
// 	fmt.Println(p)
// }



package main
import "fmt"
func main(){
	var a int      //taking input from user
	fmt.Println("Please Enter User input:")
    fmt.Scanln(&a)
	fmt.Println("The Entered Number is:",a)

}