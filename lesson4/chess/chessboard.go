// Package chess provides ...
package chess

import (
	"strconv"

	"github.com/hneis/go/lesson4/grid"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(a Point) Point {
	return Point{
		X: p.X + a.X,
		Y: p.Y + a.Y,
	}
}

type Points []Point

type Chessboard struct {
	White [16]ChessPiece
	Black [16]ChessPiece
}

func (cb Chessboard) PointExists(p Point) bool {
	return p.X >= 0 && p.X < 8 && p.Y >= 0 && p.Y < 8
}
func (cb Chessboard) SetPoint(ch ChessPiece, p Point) (ok bool) {
	if cb.PointExists(p) {
		ch.setPoint(p.X, p.Y)
		ok = true
	}

	return
}

func (cb Chessboard) CallculateAllMovies(cp ChessPiece) (movies Points) {
	for _, p := range cp.Matrix() {
		newPoint := cp.Point().Add(p)
		if cb.PointExists(newPoint) {
			movies = append(movies, newPoint)
		}
	}

	return
}

func (cb Chessboard) DrawValidMovie(c ChessPiece, points Points) {
	grid := grid.Grid{
		Rows:    9,
		Columns: 9,
		Size:    3,
	}

	grid.DrawLineTop()
	grid.DrawRow([]byte{' ', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'})

	for row := 0; row < 8; row++ {
		grid.DrawLine()
		letter := strconv.Itoa(row + 1)
		point := c.Point()
		data := make([]byte, 8)
		movies := Filter(points, func(p Point) bool {
			return p.Y == row
		})
		for col := 7; col >= 0; col-- {
			data[col] = ' '
			if point.X == col && point.Y == row {
				data[col] = c.Letter()
			}
			for i := 0; i < len(movies); i++ {
				data[movies[i].X] = 'o'
			}
		}
		data = append([]byte{letter[0]}, data...)
		grid.DrawRow(data)
	}

	grid.DrawLine()
	grid.DrawRow([]byte{' ', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'})
	grid.DrawLineBottom()
}

func Filter(d Points, f func(Point) bool) Points {
	result := make(Points, 0)
	for _, v := range d {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
