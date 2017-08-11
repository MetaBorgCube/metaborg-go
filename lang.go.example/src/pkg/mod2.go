package pkg

import "math"

func (p Point) Dist() float64 {
     return math.Sqrt((p.X * p.X) + (p.Y * p.Y))
}
