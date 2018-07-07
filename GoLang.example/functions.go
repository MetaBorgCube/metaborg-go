package main;

import "fmt";
  
func main(){    
		start := 0;
		first(start, "Hello world");
};       

func first(x int, y string){
	fmt.Println("First function");
	fmt.Println("With param x: ", x);
	fmt.Println("With param y: ", y);
	for x < 10 {
		fmt.Println("x: ", x);
		if x > 8 {
			second(x);
		};
	
		x = x + 1;
	
	};
};

func second(num int){ 
	fmt.Println("Second function num: ",num);
};