// Constant declarations
package main;

import (
  "fmt";
);

// untyped
const (
  ub = true;
  us = "xyz";
  ui0 = 42;
  ui1 = 42;
  ui2 = 42;
  ui4 = 42;
  ui8 = 42;
  uu0 = 42;
  uu1 = 42;
  uu2 = 42;
  uu4 = 42;
  uu8 = 42;
  uup = 42;
  uf1 = 1.337;
  uf2 = 1.337;
  uc1 = 1+.337i;
  uc2 = 1+.337i;
  ur  = 'x';
);

// typed
const (
  b bool = true;
  s string = "xyz";
  i0 int = 42;
  i1 int8 = 42;
  i2 int16 = 42;
  i4 int32 = 42;
  i8 int64 = 42;
  u0 uint = 42;
  u1 uint8 = 42;
  u2 uint16 = 42;
  u4 uint32 = 42;
  u8 uint64 = 42;
  up uintptr = 42;
  f1 float32 = 1.337;
  f2 float64 = 1.337;
  c1 complex64 = 1+.337i;
  c2 complex128 = 1+.337i;
  r  rune = 'x';
);

// assigning values untyped
const v = ur;

// assigning values
const (
  //ab bool = v;
  //as string = v;
  ai0 int = v;
  ai1 int8 = v;
  ai2 int16 = v;
  ai4 int32 = v;
  ai8 int64 = v;
  au0 uint = v;
  au1 uint8 = v;
  au2 uint16 = v;
  au4 uint32 = v;
  au8 uint64 = v;
  aup uintptr = v;
  af1 float32 = v;
  af2 float64 = v;
  ac1 complex64 = v;
  ac2 complex128 = v;
  ar  rune = v;
);

// assigning values coerced types
var (
  cb = true;
  cs = "xyz";
  ci0 = 42;
  ci1 = 42;
  ci2 = 42;
  ci4 = 42;
  ci8 = 42;
  cu0 = 42;
  cu1 = 42;
  cu2 = 42;
  cu4 = 42;
  cu8 = 42;
  cup = 42;
  cf1 = 1.337;
  cf2 = 1.337;
  cc1 = 1+.337i;
  cc2 = 1+.337i;
  cr  = 'x';
);


func main() {
  fmt.Printf("Type: %T\n", ub);
  fmt.Printf("Type: %T\n", us);
  fmt.Printf("Type: %T\n", ui0);
  fmt.Printf("Type: %T\n", ui1);
  fmt.Printf("Type: %T\n", ui2);
  fmt.Printf("Type: %T\n", ui4);
  fmt.Printf("Type: %T\n", ui8);
  fmt.Printf("Type: %T\n", uu0);
  fmt.Printf("Type: %T\n", uu1);
  fmt.Printf("Type: %T\n", uu2);
  fmt.Printf("Type: %T\n", uu4);
  fmt.Printf("Type: %T\n", uu8);
  fmt.Printf("Type: %T\n", uf1);
  fmt.Printf("Type: %T\n", uf2);
  fmt.Printf("Type: %T\n", uc1);
  fmt.Printf("Type: %T\n", uc2);
  fmt.Printf("Type: %T\n", ur);

  fmt.Printf("Type: %T\n", b);
  fmt.Printf("Type: %T\n", s);
  fmt.Printf("Type: %T\n", i0);
  fmt.Printf("Type: %T\n", i1);
  fmt.Printf("Type: %T\n", i2);
  fmt.Printf("Type: %T\n", i4);
  fmt.Printf("Type: %T\n", i8);
  fmt.Printf("Type: %T\n", u0);
  fmt.Printf("Type: %T\n", u1);
  fmt.Printf("Type: %T\n", u2);
  fmt.Printf("Type: %T\n", u4);
  fmt.Printf("Type: %T\n", u8);
  fmt.Printf("Type: %T\n", f1);
  fmt.Printf("Type: %T\n", f2);
  fmt.Printf("Type: %T\n", c1);
  fmt.Printf("Type: %T\n", c2);
  fmt.Printf("Type: %T\n", r);
};