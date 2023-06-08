package minstack

import "testing"

func TestMinStack(t *testing.T) {
	m := MinStack{}
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	/*
		if m.GetMin() != -3 {
			t.Errorf("Actual %d; Expected: %d", m.GetMin(), -3)
		}
		t.Logf("GetMin() = %d", m.GetMin())
		m.Pop()
		if m.Top() != 0 {
			t.Errorf("Actual %d; Expected: %d", m.Top(), 0)
		}
		t.Logf("Top() = %d", m.Top())
		if m.GetMin() != -2 {
			t.Errorf("Actual %d; Expected: %d", m.GetMin(), -2)
		}
		t.Logf("GetMin() = %d", m.GetMin())
	*/
	for m.first != nil {
		t.Log(m.Top(), m.GetMin())
		m.Pop()
	}
}
