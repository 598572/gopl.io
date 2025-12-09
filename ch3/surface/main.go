// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

// 【Go vs Java】包级常量
// Java:  private static final int WIDTH = 600;
// Go:    const width = 600
// 注意：Go的常量可以是无类型的，编译器会根据上下文推断
const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

// 【Go vs Java】包级变量初始化
// Java:  private static double sin30 = Math.sin(angle);
// Go:    var sin30, cos30 = math.Sin(angle), math.Cos(angle)
// 注意：可以同时初始化多个变量
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

// corner 计算网格角点的2D投影坐标
// 【Go vs Java】多返回值
// Java:  Point2D corner(int i, int j) { return new Point2D(sx, sy); }
// Go:    func corner(i, j int) (float64, float64) { return sx, sy }
// 注意：Go支持多返回值，不需要创建额外的对象
func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	// 【Go vs Java】返回多个值
	// Java:  return new Point2D(sx, sy);
	// Go:    return sx, sy
	return sx, sy
}

// f 计算3D表面函数值
// 【Go vs Java】多参数同类型简写
// Java:  double f(double x, double y) { ... }
// Go:    func f(x, y float64) float64 { ... }
// 注意：相同类型的连续参数可以只写一次类型
func f(x, y float64) float64 {
	// 【Go vs Java】数学函数
	// Java:  double r = Math.hypot(x, y);
	// Go:    r := math.Hypot(x, y)
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
