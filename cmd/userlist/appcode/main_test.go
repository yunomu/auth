package appcode

import (
	"strings"
	"testing"
)

func TestAddCodes(t *testing.T) {
	split := func(s string) []string {
		return strings.Split(s, ",")
	}

	result := addCodes(split("abc,def,ghi"), split("zzz,def,ghi,xxx"))

	if len(result) != 5 {
		for i, c := range result {
			t.Logf("result[%d]=%s", i, c)
		}
		t.Error("mismatch")
	}

	m := make(map[string]struct{})
	for _, c := range result {
		m[c] = struct{}{}
	}
	for _, c := range split("abc,def,ghi,xxx,zzz") {
		if _, ok := m[c]; !ok {
			t.Errorf("%s is not contains", c)
		}
	}
}

func TestDelCodes(t *testing.T) {
	split := func(s string) []string {
		return strings.Split(s, ",")
	}

	result := delCodes(split("abc,def,ghi"), split("zzz,def,ghi,xxx"))

	if len(result) != 1 {
		for i, c := range result {
			t.Logf("result[%d]=%s", i, c)
		}
		t.Error("mismatch")
	}

	m := make(map[string]struct{})
	for _, c := range result {
		m[c] = struct{}{}
	}
	for _, c := range split("abc") {
		if _, ok := m[c]; !ok {
			t.Errorf("%s is not contains", c)
		}
	}
}
