package main;

import (
    point "pkg";
    "fmt";
);

type (
  A1 = string;
  A2 = A1;
);

type (
  B1 string;
  B2 B1;
  B3 []B1;
  B4 B3;
);

func main() {
    var x string = "Go";
    x, x = "now", "now";
    p := point.Point{1,1};
    q := p.s#X;
    d := p.s#Dist();
    return "\n";
};
