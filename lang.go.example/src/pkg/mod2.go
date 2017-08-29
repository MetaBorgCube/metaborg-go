package pkg;

import "math";

func (p Point) Dist() float64 {
     return math.Sqrt(float64((p.X * p.X) + (p.Y * p.Y)))
};
