package vector

import (
	"testing"
)

// --- Add ---

func TestSum_Valid(t *testing.T) {
	a := Vector[int]{1, 2, 3}
	b := Vector[int]{4, 5, 6}
	got := a.Add(b)
	if got == nil {
		t.Fatal("Add returned nil")
	}
	expected := []int{5, 7, 9}
	for i, want := range expected {
		if got[i] != want {
			t.Errorf("[%d] got %d, want %d", i, got[i], want)
		}
	}
}

func TestSum_LengthMismatch(t *testing.T) {
	a := Vector[int]{1, 2, 3}
	b := Vector[int]{1, 2}
	if got := a.Add(b); got != nil {
		t.Errorf("expected nil for length mismatch, got %v", got)
	}
}

func TestSum_Float(t *testing.T) {
	a := Vector[float64]{1.1, 2.2}
	b := Vector[float64]{0.9, 0.8}
	got := a.Add(b)
	if got == nil {
		t.Fatal("Add returned nil")
	}
	if got[0] < 1.99 || got[0] > 2.01 {
		t.Errorf("[0] got %v, want ~2.0", got[0])
	}
	if got[1] < 2.99 || got[1] > 3.01 {
		t.Errorf("[1] got %v, want ~3.0", got[1])
	}
}

// --- Sub ---

func TestSub_Valid(t *testing.T) {
	a := Vector[int]{10, 20, 30}
	b := Vector[int]{1, 2, 3}
	got := a.Sub(b)
	if got == nil {
		t.Fatal("Sub returned nil")
	}
	expected := []int{9, 18, 27}
	for i, want := range expected {
		if got[i] != want {
			t.Errorf("[%d] got %d, want %d", i, got[i], want)
		}
	}
}

func TestSub_LengthMismatch(t *testing.T) {
	a := Vector[int]{1, 2}
	b := Vector[int]{1, 2, 3}
	if got := a.Sub(b); got != nil {
		t.Errorf("expected nil for length mismatch, got %v", got)
	}
}

func TestSub_Negative(t *testing.T) {
	a := Vector[int]{1, 2}
	b := Vector[int]{3, 5}
	got := a.Sub(b)
	if got == nil {
		t.Fatal("Sub returned nil")
	}
	if got[0] != -2 || got[1] != -3 {
		t.Errorf("got [%d, %d], want [-2, -3]", got[0], got[1])
	}
}

// --- Scl ---

func TestScl_Valid(t *testing.T) {
	a := Vector[int]{1, 2, 3}
	got := a.Scl(4)
	if got == nil {
		t.Fatal("Scl returned nil")
	}
	expected := []int{4, 8, 12}
	for i, want := range expected {
		if got[i] != want {
			t.Errorf("[%d] got %d, want %d", i, got[i], want)
		}
	}
}

func TestScl_Zero(t *testing.T) {
	a := Vector[int]{5, 10, 15}
	got := a.Scl(0)
	if got == nil {
		t.Fatal("Scl returned nil")
	}
	for i, v := range got {
		if v != 0 {
			t.Errorf("[%d] got %d, want 0", i, v)
		}
	}
}

func TestScl_Negative(t *testing.T) {
	a := Vector[int]{2, 4, 6}
	got := a.Scl(-1)
	if got == nil {
		t.Fatal("Scl returned nil")
	}
	expected := []int{-2, -4, -6}
	for i, want := range expected {
		if got[i] != want {
			t.Errorf("[%d] got %d, want %d", i, got[i], want)
		}
	}
}

func TestScl_Float(t *testing.T) {
	a := Vector[float64]{1.0, 2.0, 3.0}
	got := a.Scl(0.5)
	if got == nil {
		t.Fatal("Scl returned nil")
	}
	expected := []float64{0.5, 1.0, 1.5}
	for i, want := range expected {
		if got[i] != want {
			t.Errorf("[%d] got %v, want %v", i, got[i], want)
		}
	}
}
