package main;

import "fmt";
  
func main(){    
		i := 0;
		j := 0;
		for i < 10 {
			for j < 10 {
				fmt.Println(i+1);
				fmt.Println(j);
				
				j = j + 1;
			};
			j = 0;
			i = i + 1;
		};
};       

                            