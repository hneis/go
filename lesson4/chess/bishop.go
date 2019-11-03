// Package chess provides ...
package chess

type Bishop struct {
	point Point
	color string
}

func NewBishop(X, Y int, color string) Bishop {
	return Bishop{
		point: Point{
			X: X,
			Y: Y,
		},
		color: color,
	}
}

func (cp Bishop) Letter() byte {
	return 'B'
}

func (cp Bishop) Point() Point {
	return cp.point
}

func (cp *Bishop) setPoint(X, Y int) {
	cp.point.X = X
	cp.point.Y = Y
}

func (cp Bishop) Matrix() Points {
	return Points{
		Point{7, 7},
		Point{6, 6},
		Point{5, 5},
		Point{4, 4},
		Point{3, 3},
		Point{2, 2},
		Point{1, 1},

		Point{-1, -1},
		Point{-2, -2},
		Point{-3, -3},
		Point{-4, -4},
		Point{-5, -5},
		Point{-6, -6},
		Point{-7, -7},

		Point{-7, 7},
		Point{-6, 6},
		Point{-5, 5},
		Point{-4, 4},
		Point{-3, 3},
		Point{-2, 2},
		Point{-1, 1},

		Point{1, -1},
		Point{2, -2},
		Point{3, -3},
		Point{4, -4},
		Point{5, -5},
		Point{6, -6},
		Point{7, -7},
	}
}

func (cp Bishop) Color() string {
	return cp.color
}
