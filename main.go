package main
import "fmt"
func main(){
	var factVal uint64 = 1
	var x int 
	var e float64 = 1
	fmt.Scan(&x)
	for i:=1; i<=x; i++{
		factVal = 1
		for j:=i; j>=1; j-- {
			factVal *= uint64(j)
		}
		e += 1/float64(factVal)
		fmt.Println(e)
	}	
}