package crypto

const (
	// y^2 = x^3 - 4
	curveArgA int64 = 0
	curveArgB int64 = -4

	// mod p
	prime int64 = 199
)

type point struct {
	x int64
	y int64
}

// G = (2,2)
var G = point{
	x: 2,
	y: 2,
}

func (p *point) GetX() int64 {
	return p.x
}

func (p *point) GetY() int64 {
	return p.y
}

func (p *point) Add(q Point) Point {
	if p == nil {
		// 方便处理边界条件
		return &point{x: q.GetX(), y: q.GetY()}
	}
	var k int64
	if p.GetX() == q.GetX() {
		// k = dy/dx mod p = (3*x^2 + a) / 2*y mod p
		k = mod((3*p.GetX()*p.GetX() + curveArgA) * invNP(2*p.GetY(), prime), prime)
	} else {
		// k = (y2 - y1) / (x2 - x1) mod p
		k = mod((p.GetY() - q.GetY()) * invNP(p.GetX() - q.GetX(), prime), prime)
	}
	// xr = k^2 - x1 - x2 mod p
	xr := mod(k*k - p.GetX() - q.GetX(), prime)
	// yr = k*(x1-xr) - y1 mod p, R是直线与椭圆曲线的交点关于x轴的对称点
	yr := mod(k * (p.GetX() - xr) - p.GetY(), prime)
	return &point{x: xr, y: yr}
}

func (p *point) Multiply(k int64) Point {
	if k <= 0 {
		panic("multiply k not greater than 0")
	}
	var result Point
	var base Point = &point{x:p.x, y:p.y}
	for ; k != 0; k >>= 1{
		if k&1 !=0 {
			if result == nil {
				result = &point{x: base.GetX(), y: base.GetY()}
			} else {
				result = result.Add(base)
			}
		}
		base = base.Add(base)
	}
	return result
}
