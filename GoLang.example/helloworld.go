package main;

import "fmt";
  
func main(){    
		x := 3;
		z := true;
		fmt.Println(x);
		y := "Hello world";
		fmt.Println(y);   
		fmt.Println("yo");
		fmt.Println(4); 
		if z {
			fmt.Println("IK BEN EEN IF");
		};
		for z {
			fmt.Println("FOR");
			z = false;
		};
		test();
	
};       

func test(){
	x := 4;
	fmt.Println(x);
	y := "Hello world 2";
	fmt.Println(y);   
	fmt.Println("yo");
	fmt.Println(4); 
};

                            