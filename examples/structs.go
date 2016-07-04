package test12


type Node struct {
	value int;
	child Node
}


func main() {

	var a = Node{value: 12, child: Node{1}};
	
	println(a);
	println(a.value);
	println(a.child.value);
	
	var b = Node{1, Node{42}};
	
	println(b.value);
	println(b.child.value);
}
