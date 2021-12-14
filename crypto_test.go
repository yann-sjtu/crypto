package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvNP(t *testing.T) {
	assert.Equal(t, int64(15), invNP(-3, 23))
}

// P = (2, 2) => 103P = (96, 66)
// Q = (192, 161) => 103Q = (190, 62)

func TestPoint_Add(t *testing.T) {
	var a Point = (*point)(nil)
	for i:=0; i<103;i++ {
		a = a.Add(&point{x:2, y:2})
	}

	assert.Equal(t, int64(96), a.GetX())
	assert.Equal(t, int64(66), a.GetY())

	var b Point = (*point)(nil)
	for i:=0;i<103;i++ {
		b = b.Add(&point{x:192, y:161})
	}
	assert.Equal(t, int64(190), b.GetX())
	assert.Equal(t, int64(62), b.GetY())
}

func TestPoint_Multiply(t *testing.T) {
	var a Point = &point{x:2, y:2}
	a = a.Multiply(103)
	assert.Equal(t, int64(96), a.GetX())
	assert.Equal(t, int64(66), a.GetY())

	var b Point = &point{x:192, y:161}
	b = b.Multiply(103)
	assert.Equal(t, int64(190), b.GetX())
	assert.Equal(t, int64(62), b.GetY())
}