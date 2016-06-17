package test

import . "test"

func main(x int) int {
	var a, b int;
	
	a, b = 1, 2;
	
	a, b = func() (int, int) {
		return 3, 4;
	};
	
	println(a);
	println(b);
}