package handler

import (
	"testing"

	"github.com/yunomu/auth/lib/whitelist"
)

func TestContainEmail(t *testing.T) {
	h := NewHandler([]*whitelist.User{
		{Name: "ghi"},
		{Name: "def"},
		{Name: "abc"},
	})

	if !h.containsEmail("abc") {
		t.Errorf("error")
		t.Logf("list: %#v", h.whiteList)
	}
}
