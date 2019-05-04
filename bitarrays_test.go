package bitarray

import (
	"fmt"
	"testing"
)

func Example() {
	// New array with all values false
	a := New2D(2, 3)
	fmt.Println("Init:")
	fmt.Println(a)
	fmt.Println()

	// Set some values to true
	a.Set(0, 0, true)
	a.Set(1, 2, true)
	fmt.Println("Set some values:")
	fmt.Println(a)

	// Output:
	// Init:
	// 0 0 0
	// 0 0 0
	//
	// Set some values:
	// 1 0 0
	// 0 0 1
}

func TestIndexes(t *testing.T) {
	a := New2D(12, 20)

	var index, bit int

	// first byte
	index, bit = a.indexes(0, 0)
	if index != 0 || bit != 0 {
		t.Errorf("Wrong indexes for cell (0,0): (%v,%v) instead (0, 0)", index, bit)
	}
	index, bit = a.indexes(0, 4)
	if index != 0 || bit != 4 {
		t.Errorf("Wrong indexes for cell (0,4): (%v,%v) instead (0, 4)", index, bit)
	}
	// second byte
	index, bit = a.indexes(0, 10)
	if index != 1 || bit != 2 {
		t.Errorf("Wrong indexes for cell (0,10): (%v,%v) instead (1, 2)", index, bit)
	}
	// third byte
	index, bit = a.indexes(1, 0)
	if index != 2 || bit != 4 {
		t.Errorf("Wrong indexes for cell (1,0): (%v,%v) instead (2, 4)", index, bit)
	}
}

func BenchmarkBitGet(b *testing.B) {
	n := 500
	m := 100
	j := 250
	k := 50
	a := New2D(n, m)

	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		_ = a.Get(j, k)
	}
}

func BenchmarkBoolGet(b *testing.B) {
	n := 500
	m := 100
	j := 250
	k := 50
	a := make([]bool, m*n)
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		_ = a[j*m+k]
	}
}

func BenchmarkBitSet(b *testing.B) {
	n := 500
	m := 100
	j := 250
	k := 50
	a := New2D(n, m)
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		a.Set(j, k, true)
	}
}

func BenchmarkBoolSet(b *testing.B) {
	n := 500
	m := 100
	j := 250
	k := 50
	a := make([]bool, m*n)
	b.ResetTimer()
	for i := 0; i != b.N; i++ {
		a[j*m+k] = true
	}
}
