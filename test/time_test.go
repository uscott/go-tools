package test

import (
	"testing"
	"time"

	"github.com/uscott/go-tools/tm"
)

func TestMilliseconds(t *testing.T) {
	now := tm.UTC()
	ms, err := tm.Milliseconds(now)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Now: %+v\n", now.Format(time.StampNano))
	t.Logf("Milliseconds:  %d\n", ms)
	ms = tm.Milliseconds2(now)
	t.Logf("Milliseconds2: %d\n", ms)
}
