package main

func main(x int) {
	var a = [3]int{1, 2, 3};
	//println(a[0],a[1],a[2]);
	a[0] = "Hello";
	a[1] = "World";
	println(a[0], a[1]);
	println(a);

}
