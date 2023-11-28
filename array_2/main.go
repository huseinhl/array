package main 

import "fmt"

func main(){
		var n int 

	fmt.Scan(&n)

	a := make([]int, n)

	var x int
	res := true

	for i:=0; i<n; i++{
		fmt.Scan(&a[i])
	}
	fmt.Scan(&x)


	for i := 0; i<n; i++{
		if x == a[i] {
			res = true
		}
	}		
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}