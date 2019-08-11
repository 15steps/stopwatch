package stopwatch

import (
	// "github.com/oknotok97/stopwatch"
	"testing"
	"time"
)

func TestStopwatchInitialization(t *testing.T) {
	sw := CreateStarted()
	ticker := time.NewTicker(100 * time.Millisecond)
	<-ticker.C
	sw.Stop()
	t.Logf("Final time is %2.f", sw.Time)
	if sw.Time == 0 {
		t.Fatal("Final time is zero")
	}
}

func TestMultipleStarts(t *testing.T) {
	sw := CreateStarted()
	sw.Stop()
	sw.Start()
	if sw.Time != 0 {
		t.Fatal("Time wasn't properly reset")
	}
}

func TestStopWatch_GetElapsedTime(t *testing.T) {
	sw := CreateStarted()
	elapsed1 := sw.GetElapsedTime()
	elapsed2 := sw.GetElapsedTime()
	if elapsed1 >= elapsed2 {
		t.Fatal("Stopwatch failed to get current elapsed time")
	}
}

func TestStopWatch_GetElapsedAndReset(t *testing.T) {
	sw := CreateStarted()
	_ = sw.GetElapsedAndReset()
	if sw.GetElapsedTime() != 0 {
		t.Fatal("Stopwatch failed to reset")
	}
}