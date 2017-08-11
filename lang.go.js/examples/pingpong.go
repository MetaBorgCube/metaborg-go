package main

func main() int{
	c := make(chan string);
	go receivepong(c);
	go sendpong(c);
	println("Finished Main");
}

func receivepong(c chan string) {
	received := <-c;
	println("Received on channel C");
	println(received);
	c <- "Ping";
}

func sendpong(c chan string) {
	println("Send on channel C");
	c <- "Pong";
	received := <-c;
	println("Received");
	println(received);
}