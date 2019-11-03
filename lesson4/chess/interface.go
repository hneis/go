// Package chess provides ...
package chess

type ChessPiece interface {
	Color() string
	Matrix() Points
	Point() Point
	setPoint(X, Y int)
	Letter() byte
}
