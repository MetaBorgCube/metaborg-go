// Constant declarations
package main;

import (
  "fmt";
  "reflect";
);

type (
  A1 string;
  A2 A1;
  A3 = string;
  A4 = A3;
);

var s = "1";

var (
  a1 A1;// = s;
  a2 A2;// = s;
  a3 A3 = s;
  a4 A4 = s;
);

func main() {
  // actual type
  fmt.Printf("%T\n", a1); // main.A1
  fmt.Printf("%T\n", a2); // main.A2
  fmt.Printf("%T\n", a3); // string
  fmt.Printf("%T\n", a4); // string

  // underlying type
  fmt.Println(reflect.TypeOf(a1).Kind());  // string
  fmt.Println(reflect.TypeOf(a2).Kind());  // string
  fmt.Println(reflect.TypeOf(a3).Kind());  // string
  fmt.Println(reflect.TypeOf(a4).Kind());  // string
};
