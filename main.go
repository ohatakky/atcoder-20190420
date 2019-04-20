package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (b <= a && b <= c && a >= c) || (b >= a && b >= c && a <= c) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
