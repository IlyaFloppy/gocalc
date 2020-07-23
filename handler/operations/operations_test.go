package operations

import (
	"math"
	"testing"
)

const (
	eps = 1e-8
)

func TestAdd(t *testing.T) {
	add := Add{}
	got := add.Perform(3,5)
	expected := 8.0
	if math.Abs(got - expected) > eps {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	sub := Sub{}
	got := sub.Perform(3,9)
	expected := -6.0
	if math.Abs(got - expected) > eps {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	mul := Mul{}
	got := mul.Perform(2,7)
	expected := 14.0
	if math.Abs(got - expected) > eps {
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	div := Div{}
	got := div.Perform(1,5)
	expected := 0.2
	if math.Abs(got - expected) > eps {
		t.Fail()
	}
}