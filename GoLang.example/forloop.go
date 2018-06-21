package main;

import "fmt";
  
func main(){    
		i := 0;
		j := 0;
		for i < 10 {
			j = i;
			for j < 10 {
				fmt.Print("*");
				j = j + 1;
			};
			fmt.Println("");
			i = i + 1;
		};
};       
                           
                            