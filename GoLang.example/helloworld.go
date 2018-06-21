package main;

import "fmt";
  
func main(){    
		z := 4;
		z = z + 6;
		y := false;
		y = z < 10;
		fmt.Println(z);
		if y {
			fmt.Println("z smaller than 10");
		}else{
			fmt.Println("z greater than 10");
		};
		i := 0;
		j := 0;
		for i < 10 {
			for j < 10 {
			fmt.Println(i);
			fmt.Println(j);
			
			j = j + 1;
			};
			j = 0;
			i = i + 1;
		};
};       

func test(){
	x := 4;
	fmt.Println(x);
	y := "Hello world 2";
	fmt.Println(y);   
	fmt.Println("yo");
	fmt.Println(4); 
};

                            