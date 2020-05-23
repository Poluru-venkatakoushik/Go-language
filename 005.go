package main
import "fmt"
func main(){
	/* Syntax :
			for (a=10;a<5;a++){
				statements
			}
			 or 
			for (a<10){
				statements
			}
	*/
	
	for a:=10;a>0;a--{
		fmt.Println(a)
		
	}		 
	fmt.Println("BOOM")
	//THere's no WHILE in GO     For is go's while
	i:=10
	for i>0{
		fmt.Println(i)
		i--
	}

}