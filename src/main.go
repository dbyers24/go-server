package main

import "fmt"

func main() {
	fmt.Println("this is a meissedfasdfasdfasdfasdf")
	foo()
	fmt.Println("after foo")

	for i := 0; i < 100; i++ {
		fmt.Println(i)
			if i % 2 == 0 {
				fmt.Println("this is an even number")
			}
	}
}

func foo() {
	fmt.Println("THIS IS THE FOO MANG")
}
// control flow