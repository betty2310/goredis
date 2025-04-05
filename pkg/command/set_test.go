package command

import (
	"testing"
	"time"
)

func TestSetWithExpInSecond(t *testing.T) {
	args := []string{"set", "name", "huynh", "EXP", "5s"}
	now := time.Now()
	key, value, err := Set(args)
	if err != nil {
		t.Errorf(`args %v should return true, not err with message %v`, args, err.Error())
	}
	if key != "name" {
		t.Errorf(`args %v should have key %v`, args, "name")
	}
	if value.Value.Get() != "huynh" {
		t.Errorf(`args %v should have vale %v`, args, "huynh")
	}
	second, err := time.ParseDuration("5s")
	if err != nil {
		return
	}
	now = now.Add(second)
	if now.Compare(value.Exp) > 0 {
		t.Errorf(`the exp should be equal or larger: now: %v - value: %v`, now, value.Exp)
	}
}
