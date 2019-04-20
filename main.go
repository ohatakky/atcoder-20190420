package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	var k int
	fmt.Scan(&k)

	ts := string(s[k-1])

	str := make([]string, n)

	for i, v := range s {
		if string(v) != ts {
			str[i] = "*"
		} else {
			str[i] = string(v)
		}
	}

	tstr := strings.Join(str, "")

	fmt.Println(tstr)

}
