package main;

import "fmt";
  
func main(){    
	var a int;
	a = 3;
	fmt.Println("a: ", a);
	
	var b int = 4;
	fmt.Println("b: ", b);
	
	c,d,e := 2,"test",true;
	fmt.Println("c: ", c);
	fmt.Println("d: " , d);
	fmt.Println("e: ", e);
	
	c,d = 4,"test2";
	fmt.Println("c: ", c);
	fmt.Println("d: " , d);
	
	e,c = false,11;
	fmt.Println("c: ", c);
	fmt.Println("e: " , e);

};       

                            