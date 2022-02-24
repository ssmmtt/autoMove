package main

// upx 压缩：upx -9 -o automove.exe .\go_build_autoMove_go.exe
import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"
)

// 判断是否是工作时间
func workTime() bool {
	h := time.Now().Hour()
	if 9 < h && h < 20 {
		return true
	}
	return false
}

func main() {
	w, h := robotgo.GetScreenSize()
	fmt.Println("获取屏幕尺寸：", w, h)
	for {
		if workTime() {
			x, y := rand.Intn(w), rand.Intn(h)
			fmt.Println("移动到：", x, y)
			robotgo.MoveMouseSmooth(x, y, 0.01, 0.02)
		}
		time.Sleep(time.Minute * 5)
	}
}
