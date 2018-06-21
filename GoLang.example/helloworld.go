package main;

import "fmt";
  
func main(){    
		x := 3;
		fmt.Println(x);
		y := "Hello world";
		fmt.Println(y);   
		fmt.Println("yo");
		fmt.Println(4); 
		test();
		bools();
};       

func test(){
	x := 4;
	x = x + 4;
	fmt.Println(x);
	y := "Hello world 2";
	fmt.Println(y);   
	fmt.Println("yo");
	fmt.Println(4); 
};

func bools(){
	z := true;
	if false {
		fmt.Println("TRUE");
	};
};
                            