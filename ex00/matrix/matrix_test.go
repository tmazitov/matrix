package matrix

import (
	"testing"
)

// --- Constructor ---

func TestNewMatrix_Valid(t *testing.T) {
	cases := []struct {
		name   string
		values [][]int
		size   [2]int
	}{
		{"1x1", [][]int{{42}}, [2]int{1, 1}},
		{"2x2", [][]int{{1, 2}, {3, 4}}, [2]int{2, 2}},
		{"3x2", [][]int{{1, 2}, {3, 4}, {5, 6}}, [2]int{3, 2}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := NewMatrix(tc.values)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if m.Size() != tc.size {
				t.Errorf("got size %v, want %v", m.Size(), tc.size)
			}
		})
	}
}

func TestNewMatrix_Invalid(t *testing.T) {
	cases := []struct {
		name    string
		values  [][]int
		wantErr error
	}{
		{"nil slice", nil, ErrMatrixNil},
		{"empty outer", [][]int{}, ErrMatrixNil},
		{"empty inner", [][]int{{}}, ErrMatrixNil},
		{"jagged rows", [][]int{{1, 2}, {3}}, ErrMatrixInvalidColumn},
		{"first row longer", [][]int{{1, 2, 3}, {4, 5}}, ErrMatrixInvalidColumn},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewMatrix(tc.values)
			if err != tc.wantErr {
				t.Errorf("got err %v, want %v", err, tc.wantErr)
			}
		})
	}
}

// --- GetValue ---

func TestGetValue(t *testing.T) {
	m, _ := NewMatrix([][]int{{1, 2, 3}, {4, 5, 6}})

	cases := []struct {
		i, j   int
		want   int
		wantOk bool
	}{
		{0, 0, 1, true},
		{1, 2, 6, true},
		{-1, 0, 0, false},
		{0, -1, 0, false},
		{2, 0, 0, false},
		{0, 3, 0, false},
	}

	for _, tc := range cases {
		v, ok := m.GetValue(tc.i, tc.j)
		if ok != tc.wantOk || v != tc.want {
			t.Errorf("GetValue(%d,%d) = (%v,%v), want (%v,%v)", tc.i, tc.j, v, ok, tc.want, tc.wantOk)
		}
	}
}

// --- Add ---

func TestAdd_Valid(t *testing.T) {
	a, _ := NewMatrix([][]int{{1, 2}, {3, 4}})
	b, _ := NewMatrix([][]int{{5, 6}, {7, 8}})
	got := a.Add(b)
	if got == nil {
		t.Fatal("Add returned nil")
	}
	expected := [][]int{{6, 8}, {10, 12}}
	for i, row := range expected {
		for j, want := range row {
			v, _ := got.GetValue(i, j)
			if v != want {
				t.Errorf("[%d][%d] got %d, want %d", i, j, v, want)
			}
		}
	}
}

func TestAdd_SizeMismatch(t *testing.T) {
	a, _ := NewMatrix([][]int{{1, 2}, {3, 4}})
	b, _ := NewMatrix([][]int{{1, 2, 3}})
	if got := a.Add(b); got != nil {
		t.Errorf("expected nil for size mismatch, got %v", got)
	}
}

func TestAdd_Float(t *testing.T) {
	a, _ := NewMatrix([][]float64{{1.5, 2.5}})
	b, _ := NewMatrix([][]float64{{0.5, 1.5}})
	got := a.Add(b)
	if got == nil {
		t.Fatal("Add returned nil")
	}
	v, _ := got.GetValue(0, 0)
	if v != 2.0 {
		t.Errorf("got %v, want 2.0", v)
	}
}

// --- Sub ---

func TestSub_Valid(t *testing.T) {
	a, _ := NewMatrix([][]int{{10, 20}, {30, 40}})
	b, _ := NewMatrix([][]int{{1, 2}, {3, 4}})
	got := a.Sub(b)
	if got == nil {
		t.Fatal("Sub returned nil")
	}
	expected := [][]int{{9, 18}, {27, 36}}
	for i, row := range expected {
		for j, want := range row {
			v, _ := got.GetValue(i, j)
			if v != want {
				t.Errorf("[%d][%d] got %d, want %d", i, j, v, want)
			}
		}
	}
}

func TestSub_SizeMismatch(t *testing.T) {
	a, _ := NewMatrix([][]int{{1, 2}})
	b, _ := NewMatrix([][]int{{1}, {2}})
	if got := a.Sub(b); got != nil {
		t.Errorf("expected nil for size mismatch, got %v", got)
	}
}

func TestSub_Negative(t *testing.T) {
	a, _ := NewMatrix([][]int{{1, 2}})
	b, _ := NewMatrix([][]int{{3, 5}})
	got := a.Sub(b)
	if got == nil {
		t.Fatal("Sub returned nil")
	}
	v0, _ := got.GetValue(0, 0)
	v1, _ := got.GetValue(0, 1)
	if v0 != -2 || v1 != -3 {
		t.Errorf("got [%d, %d], want [-2, -3]", v0, v1)
	}
}

// --- Scl ---

func TestScl_Valid(t *testing.T) {
	a, _ := NewMatrix([][]int{{1, 2}, {3, 4}})
	got := a.Scl(3)
	if got == nil {
		t.Fatal("Scl returned nil")
	}
	expected := [][]int{{3, 6}, {9, 12}}
	for i, row := range expected {
		for j, want := range row {
			v, _ := got.GetValue(i, j)
			if v != want {
				t.Errorf("[%d][%d] got %d, want %d", i, j, v, want)
			}
		}
	}
}

func TestScl_Zero(t *testing.T) {
	a, _ := NewMatrix([][]int{{5, 10}, {15, 20}})
	got := a.Scl(0)
	if got == nil {
		t.Fatal("Scl returned nil")
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			v, _ := got.GetValue(i, j)
			if v != 0 {
				t.Errorf("[%d][%d] got %d, want 0", i, j, v)
			}
		}
	}
}

func TestScl_Negative(t *testing.T) {
	a, _ := NewMatrix([][]int{{2, 4}})
	got := a.Scl(-1)
	if got == nil {
		t.Fatal("Scl returned nil")
	}
	v0, _ := got.GetValue(0, 0)
	v1, _ := got.GetValue(0, 1)
	if v0 != -2 || v1 != -4 {
		t.Errorf("got [%d, %d], want [-2, -4]", v0, v1)
	}
}
