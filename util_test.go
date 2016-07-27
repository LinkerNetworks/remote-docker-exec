package main

import (
	"strings"
	"testing"
)

func TestGetEnv(t *testing.T) {
	path := getEnv("PATH")
	if len(strings.TrimSpace(path)) == 0 {
		t.Failed()
	}
}
