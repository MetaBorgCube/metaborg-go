package test

import . "test"

func a(x int) int {
	var i int;
	var x = 1;
	var y = 2;
	y = 3;
	println("a");
}

func b(y int) int {
	println(a(1));
	println("test b1");
}

func c() int {
	println(a(1));
	println("test c2");
}