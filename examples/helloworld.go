package test

import . "test"

func main(x int) int {
	// Const declarations
	const a int;
	const b = 1;
	const c, d int;
	const e, f = 1, 2;
	
	// Var declarations
	var p int;
	var q = 1;
	var r, s int;
	var t, u = 1, 2;
	
	// Assignments
	v = 2;
	
	// Call
	y = x(v);
	
	x, y = 1, 2;
	
	// Anonymous functions
	var identity = func(x int) int {
		return 1, 2;
	};
	
	// Anonymous functions with application
	var applied = func(x int) int {
		return x;
	}(1);
	
	// Goroutine with anonymous function
	go func(x int) {
        println(x);
    }("Going");
}

func x(y int) int {
	println("Function x called");
	return 1;
}