package main

import (
	"testing"
)

func TestParsing(t *testing.T) {
	time, place, err := ParseSubscribe("09:00 Moscow")
	t.Log(time, place, err)
}
