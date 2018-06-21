package main;

import "fmt";
  
func main(){ 
		j,i := 6,6;   
		y := false;
		y = j < i;
		
		if y {
			fmt.Println("j smaller than i" );
		}else{
			fmt.Println("i smaller than j");
		};
		y = j > i;
		if y {
			fmt.Println("j greater than i" );
		}else{
			fmt.Println("i greater than j");
		};	
		y = j == i;
		if y {
			fmt.Println("i equals j" );
		}else{
			fmt.Println("i not equals j");
		};
		y = !y;
		if y {
			fmt.Println("i not equals j" );
		}else{
			fmt.Println("i not not equals j");
		};

};       
         