// Package chess provides ...
package chess

type Knight struct {
	point Point
	color string
}

func NewKnight(X, Y int, color string) Knight {
	return Knight{
		point: Point{
			X: X,
			Y: Y,
		},
		color: color,
	}
}

func (cp Knight) Letter() byte {
	return 'H'
}

func (cp Knight) Point() Point {
	return cp.point
}

func (cp *Knight) setPoint(X, Y int) {
	cp.point.X = X
	cp.point.Y = Y
}

func (cp Knight) Matrix() Points {
	return Points{
		Point{1, 2},
		Point{-1, 2},
		Point{-2, 1},
		Point{-2, -1},
		Point{-1, -2},
		Point{1, -2},
		Point{2, -1},
		Point{2, 1},
	}
}

func (cp Knight) Color() string {
	return cp.color
}
