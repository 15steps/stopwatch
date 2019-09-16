package stopwatch

import (
	"math"
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

func TestStopWatch_Lap(t *testing.T) {
	sw := New()
	for i := 0; i < 10E3; i++ {
		sw.Start()
		sw.LapAndStop()
	}
	mean := 0.0
	for _, t := range sw.Laps {
		mean += t
	}
	mean /= float64(len(sw.Laps))
	t.Logf("mean: %.5f", mean)

	sd := 0.0
	for _, t := range sw.Laps {
		sd += math.Pow(t-mean, 2)
	}
	sd /= float64(len(sw.Laps))
	sd = math.Sqrt(sd)
	t.Logf("standard_deviation: %.5f", sd)
}