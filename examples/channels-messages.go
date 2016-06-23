package main

func main() {
	println("Before channels()");
	channels(); // This should be blocking and "ping" should be printed before "After channels()"
	println("After channels()");
}

func channels() {
	var messages = make(chan string);
	
	go func() {
		messages <- "ping";
	}();
	
	var msg = <-messages;
	
	println(msg);
}