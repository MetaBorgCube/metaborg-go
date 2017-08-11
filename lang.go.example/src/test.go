package main

import (
    "pkg"
    "fmt"
)

func main() {
     p := pkg.Point{1,1}
     d := p.Dist()
     fmt.Println(d)
}
