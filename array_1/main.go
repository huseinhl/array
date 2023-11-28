package main

import "fmt"

func main(){
	var n int 

	fmt.Scan(&n)

	a := make([]int, n)

	var x int
	var res int

	for i:=0; i<n; i++{
		fmt.Scan(&a[i])
	}
	fmt.Scan(&x)

	for i := 0; i<n; i++{
		if x == a[i] {
			res++
		}
	}
	fmt.Println(res)
}
