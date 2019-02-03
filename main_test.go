package main

import (
	"testing"

	"github.com/fmi/go-homework/geom"
)

func TestSampleSimpleOperations(t *testing.T) {
	a, b, c := geom.NewVector(-1, -1, 0), geom.NewVector(1, -1, 0), geom.NewVector(0, 1, 0)
	tr := NewTriangle(a, b, c)
	ray := geom.NewRay(geom.NewVector(0, 0, -1), geom.NewVector(0, 0, 1))

	if !tr.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect triangle %#v but it did not.", ray, tr)
	}
}

func TestTriangle_Intersect(t *testing.T) {
	a, b, c := geom.NewVector(5, 5, 5), geom.NewVector(10, 15, 4), geom.NewVector(15, 5, 3)
	tr := NewTriangle(a, b, c)

	ray := geom.NewRay(geom.NewVector(9, 5, -5), geom.NewVector(0.1, 0.1, 0.8))

	if !tr.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect triangle %#v but it did not.", ray, tr)
	}
}

func TestTriangle_Intersect2(t *testing.T) {
	a, b, c := geom.NewVector(-2, 2, 0), geom.NewVector(0, 3, 2), geom.NewVector(2, 2, 0)
	tr := NewTriangle(a, b, c)

	ray := geom.NewRay(geom.NewVector(0, 1, 0), geom.NewVector(0, 0.894427190999916, 0.4472135954999579))

	if !tr.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect triangle %#v but it did not.", ray, tr)
	}
}

func TestSphere_Intersect(t *testing.T) {
	a := geom.NewVector(0, 0, 0)
	s := NewSphere(a, 1)

	ray := geom.NewRay(geom.NewVector(0, 0, 4), geom.NewVector(0, 0, -1))

	if !s.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect triangle %#v but it did not.", ray, s)
	}
}

func TestSphere_Intersect2(t *testing.T) {
	a := geom.NewVector(50, 100, 50)
	s := NewSphere(a, 5)

	ray := geom.NewRay(geom.NewVector(50, 50, 50), geom.NewVector(50, 40, 50))

	if s.Intersect(ray) {
		t.Errorf("Expected ray %#v to intersect triangle %#v but it did not.", ray, s)
	}
}

func TestSampleIntersectableImplementations(t *testing.T) {
	var prim geom.Intersectable

	a, b, c, d := geom.NewVector(-1, -1, 0),
		geom.NewVector(1, -1, 0),
		geom.NewVector(0, 1, 0),
		geom.NewVector(1, 1, 0)

	prim = NewTriangle(a, b, c)
	prim = NewQuad(a, b, c, d)
	prim = NewSphere(a, 5)

	_ = prim
}
