// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// main 生成Mandelbrot分形图像并输出为PNG
func main() {
	// 【Go vs Java】多变量常量声明
	// Java:  final double xmin = -2, ymin = -2, xmax = 2, ymax = 2;
	// Go:    const (xmin, ymin, xmax, ymax = -2, -2, +2, +2)
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	// 【Go vs Java】创建图像
	// Java:  BufferedImage img = new BufferedImage(width, height, TYPE_INT_ARGB);
	// Go:    img := image.NewRGBA(image.Rect(0, 0, width, height))
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 双层循环遍历每个像素
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			// 【Go vs Java】复数类型
			// Java:  无内置复数类型，需要自定义类或使用Apache Commons Math
			// Go:    complex128 是内置的复数类型（实部和虚部都是float64）
			// 注意：complex(real, imag) 创建复数，还有complex64类型
			z := complex(x, y)

			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	// 【Go vs Java】编码PNG
	// Java:  ImageIO.write(img, "PNG", System.out);
	// Go:    png.Encode(os.Stdout, img)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// mandelbrot 计算Mandelbrot集合的颜色
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	// 【Go vs Java】复数变量
	// Java:  Complex v = new Complex(0, 0);
	// Go:    var v complex128 (零值是0+0i)
	var v complex128

	// 【Go vs Java】uint8类型
	// Java:  for (byte n = 0; n < iterations; n++)
	// Go:    for n := uint8(0); n < iterations; n++
	// 注意：uint8是无符号8位整数（0-255），等价于byte
	for n := uint8(0); n < iterations; n++ {
		// 【Go vs Java】复数运算
		// Java:  v = v.multiply(v).add(z);
		// Go:    v = v*v + z
		// 注意：Go的复数支持直接的算术运算符
		v = v*v + z

		// 【Go vs Java】复数绝对值
		// Java:  if (v.abs() > 2) { ... }
		// Go:    if cmplx.Abs(v) > 2 { ... }
		// 注意：cmplx包提供复数数学函数
		if cmplx.Abs(v) > 2 {
			// 【Go vs Java】结构体字面量
			// Java:  return new Color(255 - contrast * n);
			// Go:    return color.Gray{255 - contrast*n}
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
