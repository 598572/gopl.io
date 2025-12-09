// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

// 【Go vs Java】包级变量
// Java:  private static final Color[] palette = {Color.WHITE, Color.BLACK};
// Go:    var palette = []color.Color{color.White, color.Black}
// 注意：Go的包级变量用var声明，[]Type{...}是切片字面量
var palette = []color.Color{color.White, color.Black}

// 【Go vs Java】常量声明
// Java:  private static final int WHITE_INDEX = 0;
// Go:    const whiteIndex = 0
// 注意：Go的const必须是编译时常量，可以用括号组合多个常量
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// main 生成Lissajous图形的GIF动画
func main() {
	//!-main
	// 【Go vs Java】随机数种子
	// Java:  Random rand = new Random(System.currentTimeMillis());
	// Go:    rand.Seed(time.Now().UTC().UnixNano())
	// 注意：不设置种子，每次运行结果相同（伪随机）
	rand.Seed(time.Now().UTC().UnixNano())

	// 【Go vs Java】命令行参数判断
	// Java:  if (args.length > 0 && args[0].equals("web"))
	// Go:    if len(os.Args) > 1 && os.Args[1] == "web"
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		// 【Go vs Java】匿名函数（Lambda）
		// Java:  BiConsumer<HttpRequest, HttpResponse> handler = (req, resp) -> {...};
		// Go:    handler := func(w http.ResponseWriter, r *http.Request) {...}
		// 注意：Go的函数是一等公民，可以赋值给变量
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}

		// 【Go vs Java】注册HTTP处理器
		// Java:  server.createContext("/", handler);
		// Go:    http.HandleFunc("/", handler)
		http.HandleFunc("/", handler)
		//!-http

		// 【Go vs Java】启动HTTP服务器
		// Java:  server.start();
		// Go:    log.Fatal(http.ListenAndServe("localhost:8000", nil))
		// 注意：ListenAndServe会阻塞，log.Fatal在出错时退出程序
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	// 【Go vs Java】标准输出
	// Java:  lissajous(System.out);
	// Go:    lissajous(os.Stdout)
	// 注意：os.Stdout实现了io.Writer接口
	lissajous(os.Stdout)
}

// lissajous 生成Lissajous图形并写入输出流
// 【Go vs Java】接口参数
// Java:  void lissajous(OutputStream out)
// Go:    func lissajous(out io.Writer)
// 注意：io.Writer是Go的核心接口，任何实现Write方法的类型都满足
func lissajous(out io.Writer) {
	// 【Go vs Java】函数内常量
	// Java:  final int cycles = 5;
	// Go:    const cycles = 5
	// 注意：可以用括号组合多个常量，更清晰
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	// 【Go vs Java】随机浮点数
	// Java:  double freq = rand.nextDouble() * 3.0;
	// Go:    freq := rand.Float64() * 3.0
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator

	// 【Go vs Java】结构体字面量
	// Java:  GIF anim = new GIF(); anim.setLoopCount(nframes);
	// Go:    anim := gif.GIF{LoopCount: nframes}
	// 注意：Go的结构体可以用{字段名: 值}初始化
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// 生成每一帧
	for i := 0; i < nframes; i++ {
		// 创建图像矩形区域
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// 绘制Lissajous曲线
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 【Go vs Java】类型转换
			// Java:  (int)(x * size + 0.5)
			// Go:    int(x*size + 0.5)
			// 注意：Go的类型转换是 Type(value)，不是 (Type)value
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1

		// 【Go vs Java】append追加元素
		// Java:  anim.delays.add(delay);
		// Go:    anim.Delay = append(anim.Delay, delay)
		// 注意：append是内置函数，返回新切片，必须赋值回去
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// 【Go vs Java】取地址操作符
	// Java:  GIF.encodeAll(out, anim);
	// Go:    gif.EncodeAll(out, &anim)
	// 注意：&取地址，传递指针避免复制大结构体
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
