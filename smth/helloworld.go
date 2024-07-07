package main 
import "fmt"
import "math"
func quadratic(a int ) int  { 
	return  a*a 
}


func incrementingText(arg ...string) string { 
	completeStr:="" 
	for _, n := range arg { 
		completeStr=completeStr + n + " "
	} 
	return completeStr
}
func main() { 
	fmt.Println("Hello, World! This is my first GO! Software. ")
	name   := "Thiago" 
	salary := 2000.00 
	city  := "São Paulo"
	height, weight  := 184, 75 
	
	fmt.Println("name: ", name )
	fmt.Println("salary: ", salary)
	fmt.Println("city: ", city)
	fmt.Println("height: ", height)
	fmt.Println("weight: ", weight)
	fmt.Println("Conditions: If Salary <=2000, the person have the goverment benefit")

	if salary <= 2000 { 
		fmt.Println("The user ", name, " have the goverment benefit")
	}

	switch salary { 
		case 2000 : fmt.Println("User is on benefit bc is on exact value ")
	} 

	for i := 0 ; i < 5 ; i ++ { 
		fmt.Println("For Loop ", i , " times")
	}

	fmt.Println("Returning the quadratic of values on the list ")
	values:=[]int{2,3,4,10,20} //var values = []int {10,20,30,40}
	for i, n := range values { 
		fmt.Println("Quadratic value : " , quadratic(n) , " Index: ", i )

	}

	fmt.Println("Testing a arg function : ", incrementingText("Hi,", "Good", "Morning!", "How", "are", "u?"))
	fmt.Println("Testing Defer:  ")
	defer fmt.Println("Defer One! ")
	defer fmt.Println("Defer two! ") //LIFO
	fmt.Println("Next step, ", math.Pi)
	



}


