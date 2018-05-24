package main;

import "fmt";

func main(){
	test3();
};

func test2(){
	fmt.Println("hello world 2");
};

func test3(){
	fmt.Println("Test");
test2();
};