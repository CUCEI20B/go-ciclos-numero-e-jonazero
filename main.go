package main
import "fmt"
func main(){
	var limite int
	var e float64 = 1
	var aux float64 = 1
	fmt.Scan(&limite)
	for i:=1; i<limite; i++{
		aux = aux/float64(i)
		e = e + aux
	}
	fmt.Println(e)
}