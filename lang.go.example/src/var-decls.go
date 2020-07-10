// Constant declarations
package main;

import (
  "fmt";
);

// untyped and initialized
var (
  vub = true;
  vus = "xyz";
  vui0 = 42;
  vui1 = 42;
  vui2 = 42;
  vui4 = 42;
  vui8 = 42;
  vuu0 = 42;
  vuu1 = 42;
  vuu2 = 42;
  vuu4 = 42;
  vuu8 = 42;
  vuup = 42;
  vuf1 = 1.337;
  vuf2 = 1.337;
  vuc1 = 1+.337i;
  vuc2 = 1+.337i;
  vur  = 'x';
);

// typed and initialized
var (
  vib bool = true;
  vis string = "xyz";
  vii0 int = 42;
  vii1 int8 = 42;
  vii2 int16 = 42;
  vii4 int32 = 42;
  vii8 int64 = 42;
  viu0 uint = 42;
  viu1 uint8 = 42;
  viu2 uint16 = 42;
  viu4 uint32 = 42;
  viu8 uint64 = 42;
  viup uintptr = 42;
  vif1 float32 = 1.337;
  vif2 float64 = 1.337;
  vic1 complex64 = 1+.337i;
  vic2 complex128 = 1+.337i;
  vir  rune = 'x';
);

// typed but uninitialized
var (
  b bool;
  s string;
  i0 int;
  i1 int8;
  i2 int16;
  i4 int32;
  i8 int64;
  u0 uint;
  u1 uint8;
  u2 uint16;
  u4 uint32;
  u8 uint64;
  up uintptr;
  f1 float32;
  f2 float64;
  c1 complex64;
  c2 complex128;
  r  rune;
);


