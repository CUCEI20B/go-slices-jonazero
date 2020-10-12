package main
import "fmt"
func main(){
	var x uint = 0
	var y uint = 0
	var z uint = 0
	fmt.Scan(&x)
	s:=make([]uint,x)
	for i:=0; i<int(x); i++{
		fmt.Scan(&y)
		s = append(s,y)
		z += y
	}
	fmt.Print(z)
}