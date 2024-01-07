package example_test

import (
	"fmt"

	"github.com/juntakoman123/go_testing/example"
)

func ExampleHello() {
	result := example.Hello()
	fmt.Println(result)
	// Output: hello
}

func ExampleUnordered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	// Unordered output:
	// 3
	// 2
	// 1
}

func ExampleOrdered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	// output:
	// 1
	// 2
	// 3
}
