package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

const (
	Height     = 6
	Width      = Height * 2
	GlassWidth = 2
	TotalSands = Height * (Height + 1)
)

type Glass struct {
	upper      [][]bool
	lower      [][]bool
	upperState []bool
	lowerState []bool
	screenH    int
	screenW    int
}

func NewGlass() *Glass {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))

	g := &Glass{
		upper:      make([][]bool, Height),
		lower:      make([][]bool, Height),
		upperState: make([]bool, Height),
		lowerState: make([]bool, Height),
		screenW:    width,
		screenH:    height,
	}

	for i := range g.upperState {
		g.upperState[i] = true
		g.lowerState[i] = false
	}

	for i := range g.upper {
		g.upper[i] = make([]bool, Width)
		g.lower[i] = make([]bool, Width)
		for j := Height; j <= Height+i; j++ {
			g.upper[i][j] = true
			g.upper[i][Width-j-1] = true
		}
	}
	return g
}

func (g *Glass) SandFlow() {
	upperIdx := -1
	// for upper:
	// true   ->  has sand
	// false  ->  no sand
	for i := Height - 1; i >= 0; i-- {
		if g.upperState[i] {
			upperIdx = i
			break
		}
	}

	for j := Height; j <= Height+upperIdx; j++ {
		if g.upper[upperIdx][j] {
			g.upper[upperIdx][j] = false
			break
		} else if g.upper[upperIdx][Width-j-1] {
			g.upper[upperIdx][Width-j-1] = false
			if Width-j-1 == Height-upperIdx-1 {
				g.upperState[upperIdx] = false
			}
			break
		}
	}

	// for lower:
	// true  ->  full
	// false ->  not full
	lowerIdx := -1
	for i := Height - 1; i >= 0; i-- {
		if !g.lowerState[i] {
			lowerIdx = i
			break
		}
	}
	for j := Height; j <= Height+lowerIdx; j++ {
		if !g.lower[lowerIdx][j] {
			g.lower[lowerIdx][j] = true
			break
		} else if !g.lower[lowerIdx][Width-j-1] {
			g.lower[lowerIdx][Width-j-1] = true
			if Width-j-1 == Height-lowerIdx-1 {
				g.lowerState[lowerIdx] = true
			}
			break
		}
	}
}

func (g *Glass) PrintGlass() {
	totalHeight := Width
	startY := (g.screenH - totalHeight) / 2
	startX := (g.screenW - Width) / 2
	var out bytes.Buffer
	for i := Height - 1; i >= 0; i-- {
		for j := 0; j < Width; j++ {
			if g.upper[i][j] {
				out.WriteRune('*')
			} else {
				out.WriteRune(' ')
			}
		}
		out.WriteRune('\n')
		fmt.Printf("\033[%d;%dH%s\n", startY+Height-i-1, startX, out.String())
		out.Reset()
	}
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if g.lower[i][j] {
				out.WriteRune('*')
			} else {
				out.WriteRune(' ')
			}
		}
		out.WriteRune('\n')
		fmt.Printf("\033[%d;%dH%s\n", startY+Height+i, startX, out.String())
		out.Reset()
	}
}

func getScreenSize() {
	w, h, _ := term.GetSize(int(os.Stdout.Fd()))
	content := "Hello, World!"

	// 计算开始打印的位置
	startX := (w - len(content)) / 2
	startY := h / 2

	// 移动光标到指定位置
	fmt.Printf("\033[%d;%dH", startY, startX)

	// 打印内容
	fmt.Println(content)
}

func main() {
	g := NewGlass()
	// for {
	// 	fmt.Print("\033[2J\033[3J\033[H")
	// 	getScreenSize()
	// }

	for i := 0; i <= TotalSands; i++ {
		// 清除屏幕并移动光标到左上角
		// fmt.Print("\033[2J\033[H")
		fmt.Print("\033[2J\033[3J\033[H")

		// 打印动画帧
		// fmt.Println("Frame", i)
		g.PrintGlass()
		// 等待一段时间
		time.Sleep(time.Second / 5)
		g.SandFlow()
	}
}
