package main

import (
	"testing"
)

func TestParsing(t *testing.T) {
	time, place, err := ParseSubscribe("09:00 Moscow")
	t.Log(time, place, err)
}
func TestSqlAdd(t *testing.T) {
	time, place, _ := ParseSubscribe("09:00 Sydney")
	testSub := subscription{
		chatId: 0,
		time:   time,
		city:   place,
	}
	err := sqlAddSubscription(testSub)
	if err != nil {
		t.Error(err)
	}

}
