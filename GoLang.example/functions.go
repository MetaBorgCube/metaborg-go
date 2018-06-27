package main;

import "fmt";
  
func main(){    
		first(2);
};       

func first(x int){
	fmt.Println("First function with param x: ", x);
	second();
};

func second(){
	fmt.Println("Second function");
};