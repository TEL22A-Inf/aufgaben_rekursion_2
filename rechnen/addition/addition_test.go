package addition

import "fmt"

func ExampleAdd() {
	fmt.Println(Add(3, 4))
	fmt.Println(Add(3, 5))
	fmt.Println(Add(2, 4))
	fmt.Println(Add(0, 4))
	fmt.Println(Add(4, 0))

	// Output:
	// 7
	// 8
	// 6
	// 4
	// 4
}
