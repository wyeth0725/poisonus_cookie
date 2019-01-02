package main

import "fmt"

func main() {
	var a,b,c int 
	fmt.Scan(&a,&b,&c)
	if a + b + 1 > c{
		fmt.Println(c + b)
	}else{
		fmt.Println(a + 2 * b + 1)
	}
}
