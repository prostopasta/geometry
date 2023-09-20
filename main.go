package main

import (
    "fmt"
    "geometry/geometry_pkg"
)

func main() {
    rect := geometry_pkg.Rectangle{Width: 10, Height: 5}
    area := rect.Area()
    scaledW, scaledH := rect.Scale(2.0)
    fmt.Println(area)
    fmt.Println(scaledW, scaledH)

}