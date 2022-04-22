package main

import (
	"testing"
	"time"
)

func TestParsing(t *testing.T) {
	time, place, err := ParseSubscribe("09:00 Moscow")
	t.Log(time, place, err)
}
func TestSqlAdd(t *testing.T) {
	testSub := subscription{
		chatId: 0,
		time:   time.Now(),
		city:   "Moscow",
	}
	err := sqlAddSubscription(testSub)
	t.Log(err)
}
