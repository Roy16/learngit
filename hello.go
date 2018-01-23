package main

import (
	"fmt"
	//"io"
)

const (
	_          = iota
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}

//fmt.Println("hello,world!")
/*
6: 0110
11:1011
--------
&  0010
|  1111 15
^  110
&^
		   	var a int = 65
		   	var c DD
		   	b := string(a)
		   	c = "中文"
		   	fmt.Println(b, c)

		   	e := true
		   	if a, b, c := 1, 2, 3; a+b+c > 6 {
		   		fmt.Println("大于6")
		   	} else {
		   		fmt.Println("小于6")
		   		fmt.Println(e)
		   	}
		   	fmt.Println(e)

		   	z := 1
		   	switch z {
		   	case 0:
		   		fmt.Println("z=0")
		   	case 1:
		   		fmt.Println("z=1")
		   	}
		   	fmt.Println(z)
		   	switch {
		   	case z >= 0:
		   		fmt.Println("z=0")
		   		fallthrough
		   	case z >= 1:
		   		fmt.Println("z=1")
		   	}

		   	//goto break continue

		   	for i := 0; i < 10; i++ {
		   		if i > 3 {
		   			break
		   		} else {
		   			fmt.Println(i)
		   		}
		   	}
		   lable:
		   	for i := 0; i < 10; i++ {
		   		for {
		   			fmt.Println(i)
		   			continue lable
		   		}
		   	}

		a := [3]int{1, 2, 3}
		fmt.Println(a)

		n := make(map[int]string)

		n[1] = "Hello"
		n[2] = "World!"
		fmt.Println(n)

	a, _, c, d := 1, 2, 3, 4
	fmt.Println(a)
	//fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
*/
