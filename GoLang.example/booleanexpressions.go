package main;

import "fmt";
  
func main(){ 
		j,i := 6,9;   
		y := false;
		y = j < i;
		
		if y {
			fmt.Println("j: ", j, " is smaller than i: ",i );
		}else{
			fmt.Println("j: " , j, " not smaller than i: ",i);
		};
		y = j > i;
		if y {
			fmt.Println("j: ", j, " is greater than i: ",i );
		}else{
			fmt.Println("j: " , j, " not greater than i:",i);
		};
		y = j == i;
		if y {
			fmt.Println("j: ", j, " equals i: ",i );
		}else{
			fmt.Println("j: " , j, " not equals i: ",i);
		};
		y = !y;
		if y {
			fmt.Println("j: ", j, " not equals i: ",i );
		}else{
			fmt.Println("j: " , j, " not not equals i:",i);
		};

};       
         