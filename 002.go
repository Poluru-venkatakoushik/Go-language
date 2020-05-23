package main
import "fmt"
func main(){
	/*variables can be decalred in many ways*/
	var a int = 100
	var b float32 = 12.5
	var c float64 = 12.5
	var(
		d = 10
		e = true
	)
	f:=false
	g:=12.5
	h:=12.
	/*
		The problem with using := method is that when we initialise the variable to 12. it is stored as float64 . to overocme this we can use var (variablename) float32=12.
	*/
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}