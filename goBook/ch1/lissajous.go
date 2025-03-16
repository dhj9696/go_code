// lissajous产生随机丽萨茹图形的GIF动画
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 //first color in palette
	blackIndex = 1 //next color in palette
)

func main() {
	//图像的顺序是确定的，除非我们播种
	//使用当前时间的伪随机数生成器。
	//感谢兰德尔·麦克弗森指出遗漏。
	//rand.Seed(time.Now().UTC().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     //x 振荡器的完整转数
		res     = 0.001 //角分辨率
		size    = 100   //图像画布覆盖范围[-尺寸..+尺寸]
		nframes = 64    //动画帧数
		delay   = 8     //以 10 毫秒为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 //y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	//if err != nil {
	//	return
	//}

	//create, _ := os.Create("out.gif")
	//lissajous(create)
}
