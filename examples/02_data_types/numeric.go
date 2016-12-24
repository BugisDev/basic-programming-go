package main

import "fmt"

func Numeric() {
	var (
		x   int
		y   int32
		z   int64
		xy  uint
		xz  uint32
		xx  uint64
		xyz float32
		zyx float64
	)

	// default is numeric 0
	fmt.Println("Default value of numeric is", x)
	x = 10
	// type conversion
	y = int32(x)
	z = int64(x)
	xy = uint(x)
	xz = uint32(x)
	xx = uint64(x)
	xyz = float32(x)
	zyx = float64(x)

	fmt.Printf("x: %d. Type: %T\n", x, x)
	fmt.Printf("y: %d. Type: %T\n", y, y)
	fmt.Printf("z: %d. Type: %T\n", z, z)
	fmt.Printf("xy: %d. Type: %T\n", xy, xy)
	fmt.Printf("xz: %d. Type: %T\n", xz, xz)
	fmt.Printf("xx: %d. Type: %T\n", xx, xx)
	fmt.Printf("xyz: %.2f. Type: %T\n", xyz, xyz)
	fmt.Printf("zyx: %.2f. Type: %T\n", zyx, zyx)
}
