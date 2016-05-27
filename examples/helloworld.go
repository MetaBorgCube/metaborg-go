package test

import . "test"

func a(x int) int {
	println();
}

func b(y int) int {
	println(a(1));
	println("test b1");
}

func c() int {
	println(a(1));
	println("test c2");
}