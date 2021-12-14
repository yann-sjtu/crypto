package crypto

type Point interface {
	GetX() int64
	GetY() int64
	Add(Point) Point
	Multiply(int64) Point
}
