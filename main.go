package main
import "fmt"
func main(){
	var e float64 = 0
	var aux float64 = 1
	var x int
	fmt.Scan(&x)
	for i:=0; i<x; i++{
		if i !=0{
		aux = aux/float64(i)
		}
		e = e + aux
	}
	fmt.Println(e)
}