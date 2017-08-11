package main

func main() {
	// Const declarations
	const a int = 0;
	const b = 1;
	const c, d int = 3, 4;
	const e, f = 1, 2;
	
	// Var declarations
	var p int;
	var q = 1;
	var r, s int;
	var t, u = 1, 2;
	
	// Assignments
	p = 1;
	r, s = 2, 3;
	r, s = func() (int, int) {
		return 4, 5;
	}();
	
	// Anonymous functions
	var identity = func(x int) int {
		return 1;
	};
	
	// Call
	p = identity(q);
	
	// Anonymous functions with application
	var v = func(x int) int {
		return x;
	}(p);
	
	// Goroutine with anonymous function
	go func(x string) {
        println(x);
    }("Going");
    
    y := 123;
    x := x(y);
    
    channels();
    
    // Log values
    println(a);
    println(b);
    println(c);
    println(d);
    println(e);
    println(f);
    println(p);
    println(q);
    println(r);
    println(s);
    println(t);
    println(u);
    println(v);
    println(x);
    println(y);
}

func x(y int) int {
	println("Function x called");
	return 1;
}

func channels() {
	var messages = make(chan string);
	
	go func() {
		messages <- "ping";
	}();
	
	var msg = <-messages;
	
	println(msg);
}