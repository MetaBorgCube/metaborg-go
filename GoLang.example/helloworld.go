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
		for i < 20 {
			fmt.Println(i);
			i = i + 1;
		};
};       


                            