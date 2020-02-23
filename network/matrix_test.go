package network

import (
	"reflect"
	"testing"
)

func TestApplyFunction(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	ApplyFunction(a, func(x float64) float64 {
		return x + 1
	})

	// Excepted value
	r := [][]float64{
		{2, 3, 4},
		{5, 6, 7},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("ApplyFunction(fn(x)=x+1) failed, excepted %v, got %v", r, a)
	}
}

func TestSum(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	Sum(a, b)

	// Excepted value
	r := [][]float64{
		{2, 4, 6},
		{8, 10, 12},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("Sum(%v) failed, excepted %v, got %v", b, r, a)
	}
}

func TestDifference(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	Difference(a, b)

	// Excepted value
	r := [][]float64{
		{0, 0, 0},
		{0, 0, 0},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("Difference(%v) failed, excepted %v, got %v", b, r, a)
	}
}

func TestDotProduct(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := [][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	// Actual value
	p := DotProduct(a, b)

	// Excepted value
	r := [][]float64{
		{58, 64},
		{139, 154},
	}

	if !reflect.DeepEqual(p, r) {
		t.Errorf("DotProduct(%v) failed, excepted %v, got %v", b, r, p)
	}
}

func TestTranspose(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	Transpose(a)

	r := [][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("Transpose failed, excepted %v, got %v", r, a)
	}
}