package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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
	upper       [][]bool
	lower       [][]bool
	upperState  []bool
	lowerState  []bool
	screenH     int
	screenW     int
	totalHeight int
}

func NewGlass() *Glass {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))

	g := &Glass{
		upper:       make([][]bool, Height),
		lower:       make([][]bool, Height),
		upperState:  make([]bool, Height),
		lowerState:  make([]bool, Height),
		screenW:     width,
		screenH:     height,
		totalHeight: Width + 4,
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

func (g *Glass) reset() {
	for i := range g.upperState {
		g.upperState[i] = true
		g.lowerState[i] = false
	}

	for i := range g.upper {
		for j := Height; j <= Height+i; j++ {
			g.upper[i][j] = true
			g.upper[i][Width-j-1] = true
		}
	}
	for i := range g.lower {
		for j := range g.lower[i] {
			g.lower[i][j] = false
		}
	}
}

func (g *Glass) sandFlow() {
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
	if g.lowerState[0] {
		g.reset()
	}
}

func (g *Glass) PrintGlass() {
	startY := (g.screenH - g.totalHeight) / 2
	startX := (g.screenW - Width - 2*GlassWidth) / 2
	fmt.Printf("\033[%d;%dH%s\n", startY, startX, strings.Repeat("=", Width+2*GlassWidth))
	var out bytes.Buffer
	for i := Height - 1; i >= 0; i-- {
		for k := 0; k < Height-i-1; k++ {
			out.WriteRune(' ')
		}
		out.WriteRune('\\')
		out.WriteRune('\\')
		s := Height - i - 1
		for j := s; j <= Width-s-1; j++ {
			if g.upper[i][j] {
				out.WriteRune('*')
			} else {
				out.WriteRune(' ')
			}
		}
		out.WriteRune('/')
		out.WriteRune('/')
		out.WriteRune('\n')
		fmt.Printf("\033[%d;%dH%s\n", startY+Height-i, startX, out.String())
		out.Reset()
	}

	fmt.Printf("\033[%d;%dH%s\n", startY+Height+1, startX, strings.Repeat(" ", Height)+"\\\\//")
	fmt.Printf("\033[%d;%dH%s\n", startY+Height+2, startX, strings.Repeat(" ", Height)+"//\\\\")

	for i := 0; i < Height; i++ {
		for k := 0; k < Height-i-1; k++ {
			out.WriteRune(' ')
		}
		out.WriteRune('/')
		out.WriteRune('/')
		s := Height - i - 1
		for j := s; j <= Width-s-1; j++ {
			if g.lower[i][j] {
				out.WriteRune('*')
			} else {
				out.WriteRune(' ')
			}
		}
		out.WriteRune('\\')
		out.WriteRune('\\')
		out.WriteRune('\n')
		fmt.Printf("\033[%d;%dH%s\n", startY+Height+i+3, startX, out.String())
		out.Reset()
	}
	fmt.Printf("\033[%d;%dH%s\n", startY+Width+3, startX, strings.Repeat("=", Width+2*GlassWidth))
	now := time.Now()

	fmt.Printf("\033[%d;%dH%s\n", startY+Width+5, (g.screenW-18)/2, now.Format("2006-01-02 15:04:05"))
}

func main() {
	g := NewGlass()
	for {
		fmt.Print("\033[2J\033[3J\033[H")
		g.PrintGlass()
		time.Sleep(time.Second / 3)
		g.sandFlow()
	}
}
