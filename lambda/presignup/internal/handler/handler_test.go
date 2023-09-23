package handler

import (
	"testing"

	"github.com/yunomu/auth/lib/userlist"
)

func TestContainEmail(t *testing.T) {
	h := NewHandler([]*userlist.User{
		{Name: "ghi"},
		{Name: "def"},
		{Name: "abc"},
	})

	if !h.containsEmail("abc") {
		t.Errorf("error")
		t.Logf("list: %#v", h.userlist)
	}
}
