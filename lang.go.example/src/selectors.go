package p;

type B struct {
  z string;

};

var c = 10;

type A struct {
  x int;
  y float32;
//  arr [c]string;
};

type T0 struct { x int; };
type T1 struct { y int; };
//type T2 struct { z int; T1; *T0; };

func (T1) M1() {};

//func (*T0) M0() {};
//func (*T2) M2() {};
//type Q *T2;
var t T0;     // with t.T0 != nil
//var p *T2;    // with p != nil and (*p).T0 != nil
//var q Q = p;
//var r = 10;
//var s = r;
var r = t.x;
