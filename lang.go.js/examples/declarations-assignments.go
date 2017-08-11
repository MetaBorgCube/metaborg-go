package examples

func main(x int) int {
	var a, b int; // Declarations without instantation, with type
	
	a, b = 1, 2; // Multiple assignments
	c, d := 5, 6; // Multiple declaration with assignments
	
	// Assignment form multiple return values
	a, b = func() (int, int) {
		return 3, 4;
	};
	
	println(a);
	println(b);
	println(c);
	println(d);
}