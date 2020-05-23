package main
import "fmt"
func main(){
	a:=10
	/*Syntax is 
		if (condition){
			statements
		}else if (condition){
			statements

		}else{
			statements
		}
	*/
	// There is no ternary operator in GO
	// if variable = (val) (var<50){st.} also works in Go
	if (a>5){
		fmt.Println("a is greater than 5")
	}else{
		fmt.Println("a is less than 5")
	}
}