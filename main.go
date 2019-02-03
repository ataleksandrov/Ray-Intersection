package main

import (
	"math"
	"sync"

	"github.com/fmi/go-homework/geom"
)

type Triangle struct {
	A, B, C geom.Vector
}

type Quad struct {
	A, B, C, D geom.Vector
}

type Sphere struct {
	Origin geom.Vector
	R      float64
}

func NewTriangle(a, b, c geom.Vector) Triangle {
	return Triangle{
		A: a,
		B: b,
		C: c,
	}
}

func NewQuad(a, b, c, d geom.Vector) Quad {
	return Quad{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func NewSphere(origin geom.Vector, r float64) Sphere {
	return Sphere{
		Origin: origin,
		R:      r,
	}
}

func (tr Triangle) Intersect(ray geom.Ray) bool {
	// a.k.a Möller–Trumbore intersection algorithm
	const EPSILON float64 = 0.0000001
	e1 := Subtract(tr.B, tr.A)
	e2 := Subtract(tr.C, tr.A)
	p := CrossProduct(ray.Direction, e2)
	det := DotProduct(e1, p)

	if det > -EPSILON && det < EPSILON {
		return false
	}

	f := 1.0 / det
	s := Subtract(ray.Origin, tr.A)
	u := f * DotProduct(s, p)

	if u < 0.0 || u > 1.0 {
		return false
	}

	q := CrossProduct(s, e1)
	v := f * DotProduct(ray.Direction, q)

	if v < 0.0 || u+v > 1.0 {
		return false
	}
	t := f * DotProduct(e2, q)

	return t > EPSILON

}

func (quad Quad) Intersect(ray geom.Ray) bool {
	// Disclaimer: doesn't work for complex quads :( (had no time to implement "Triangulating")
	var wg sync.WaitGroup
	const cnt int = 2
	ch := make(chan bool, cnt)
	defer close(ch)
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var tr Triangle
			if i == 0 {
				tr = Triangle{quad.A, quad.B, quad.C}
			} else {
				tr = Triangle{quad.A, quad.C, quad.D}
			}
			ch <- tr.Intersect(ray)
		}(i)
	}
	wg.Wait()

	return <-ch || <-ch
}

func (sph Sphere) Intersect(ray geom.Ray) bool {
	L := Subtract(ray.Origin, sph.Origin)
	a := DotProduct(ray.Direction, ray.Direction)
	b := 2 * DotProduct(ray.Direction, L)
	c := DotProduct(L, L) - math.Pow(sph.R, 2)
	const scaleNum float64 = 0.01
	closeDirectionPoint := Add(ray.Origin, Multiply(ray.Direction, scaleNum))
	return b*b-4*a*c >= 0 && Distance(sph.Origin, ray.Origin) > Distance(sph.Origin, closeDirectionPoint)
}

func Distance(a, b geom.Vector) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2))
}

func DotProduct(a, b geom.Vector) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

func CrossProduct(a, b geom.Vector) geom.Vector {
	return geom.Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func Add(a, b geom.Vector) geom.Vector {
	return geom.Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func Subtract(a, b geom.Vector) geom.Vector {
	return geom.Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func Multiply(a geom.Vector, f float64) geom.Vector {
	return geom.Vector{
		X: a.X * f,
		Y: a.Y * f,
		Z: a.Z * f,
	}
}

func main() {

}
