package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 10; i++ {
        // 清除屏幕并移动光标到左上角
        fmt.Print("\033[2J\033[H")
        // 打印动画帧
        fmt.Println("Frame", i)
        // 等待一段时间
        time.Sleep(time.Second / 2)
    }
}

