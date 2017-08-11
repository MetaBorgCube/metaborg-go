package main

func main() {
	var i = 41;

	var p = &i;         // point to i
	println(*p); // read i through the pointer
	*p = 21 ;        // set i through the pointer
	println(i);  // see the new value of i

	p = &j;         // point to j
	println(j); // see the new value of j
}