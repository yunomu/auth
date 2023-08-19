package handler

import "testing"

func TestContainEmail(t *testing.T) {
	h := NewHandler([]string{
		"ghi",
		"def",
		"abc",
	})

	if !h.containsEmail("abc") {
		t.Errorf("error")
		t.Logf("list: %#v", h.whiteList)
	}
}
