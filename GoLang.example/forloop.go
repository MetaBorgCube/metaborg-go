package main;

import "fmt";
  
func main(){    
		i := 0;
		j := 0;
		amount := 50;
		for i < amount {
			j = i;
			for j < amount {
				fmt.Print("*");
				j = j + 1;
			};
			fmt.Println("");
			i = i + 1;
		};
		
	
};       
                           
                            