package main

/*
	Output should be:
	1
	2
	3
	1
*/

func intSeq() func() int {
    i := 0;
    return func() int {
        i += 1;
        return i;
    };
}

func main() {
    nextInt := intSeq();

    println(nextInt());
    println(nextInt());
    println(nextInt());

    newInts := intSeq();
    println(newInts());
}