// Package grid provides ...
package grid

import "fmt"

type Grid struct {
	Rows    int
	Columns int
	Size    int
}

func (g Grid) DrawLineTop() {
	fmt.Printf("┌")
	for i := 0; i < g.Columns; i++ {
		for j := 0; j < g.Size; j++ {
			fmt.Printf("─")
		}
		if i != g.Columns-1 {
			fmt.Printf("┬")
		} else {
			fmt.Printf("┐")
		}
	}
	fmt.Printf("\n")
}

func (g Grid) DrawLineBottom() {
	fmt.Printf("└")
	for i := 0; i < g.Columns; i++ {
		for j := 0; j < g.Size; j++ {
			fmt.Printf("─")
		}
		if i != g.Columns-1 {
			fmt.Printf("┴")
		} else {
			fmt.Printf("┘")
		}
	}
	fmt.Printf("\n")
}

func (g Grid) DrawRow(data []byte) {
	fmt.Printf("│")
	for i := 0; i < g.Columns; i++ {
		for j := 0; j < g.Size; j++ {
			if j == 1 {
				fmt.Printf(string(data[i]))
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("│")

	}
	fmt.Printf("\n")
}

func (g Grid) DrawLine() {
	fmt.Printf("├")
	for i := 0; i < g.Columns; i++ {
		for j := 0; j < g.Size; j++ {
			fmt.Printf("─")
		}
		if i != g.Columns-1 {
			fmt.Printf("┼")
		} else {
			fmt.Printf("┤")
		}

	}
	fmt.Printf("\n")
}
