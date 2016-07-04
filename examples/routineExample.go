package main

func main() int{
	heavyOperation();			// Will block main routine
	println("Finished1");
	go heavyOperation();	// Will run concurrently
	println("Finished2");
}

func heavyOperation(){
	println(1);
	println(2);
	println(3);
}