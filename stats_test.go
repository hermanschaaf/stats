package stats

import (
	"fmt"
	"testing"
)

func TestMean(t *testing.T) {
	m := Mean([]int{1, 2, 3, 4, 5})
	if m != 3.0 {
		t.Errorf("%.1f != %.1f", m, 3.0)
	}

	m = Mean([]int{1, 2, 3, 4, 5, 6})
	if m != 3.5 {
		t.Errorf("%.1f != %.1f", m, 3.5)
	}

	m = Mean([]int{1})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}

	m = Mean([]int{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestMedian(t *testing.T) {
	m := Median([]int{5, 3, 4, 2, 1})
	if m != 3.0 {
		t.Errorf("%.1f != %.1f", m, 3.0)
	}

	m = Median([]int{6, 3, 2, 4, 5, 1})
	if m != 3.5 {
		t.Errorf("%.1f != %.1f", m, 3.5)
	}

	m = Median([]int{1})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}

	m = Median([]int{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestMode(t *testing.T) {
	m := Mode([]int{5, 3, 4, 2, 1})
	if m != 1.0 {
		t.Errorf("%.1f != %.1f", m, 1.0)
	}

	m = Mode([]int{5, 5, 5, 5, 3, 4, 2, 1})
	if m != 5.0 {
		t.Errorf("%.1f != %.1f", m, 5.0)
	}

	m = Mode([]int{5, 5, 3, 3, 3, 4, 4, 1})
	if m != 3.0 {
		t.Errorf("%.1f != %.1f", m, 3.0)
	}

	m = Mode([]int{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}

func TestStandardDeviation(t *testing.T) {
	m := StandardDeviation([]int{5, 3, 4, 2, 1})
	if fmt.Sprintf("%.1f", m) != fmt.Sprintf("%.1f", 1.4) {
		t.Errorf("%.1f != %.1f", m, 1.4)
	}

	m = StandardDeviation([]int{})
	if m != 0.0 {
		t.Errorf("%.1f != %.1f", m, 0.0)
	}
}
