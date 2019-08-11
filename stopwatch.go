package stopwatch

import (
	"time"
)

// StopWatch struct that represents a StopWatch instance
type StopWatch struct {
	InitialTime time.Time
	Time        float64
}

// CreateStarted create and start new StopWatch
func CreateStarted() *StopWatch {
	init := time.Now().UTC()
	return &StopWatch{
		InitialTime: init,
		Time:        0,
	}
}
// New creates new instance of StopWatch
func New() *StopWatch {
	return &StopWatch{}
}

// Start starts StopWatch
func (s *StopWatch) Start() {
	s.InitialTime = time.Now().UTC()
	s.Time = 0
}

// Stop stops ongoing StopWatch
func (s *StopWatch) Stop() {
	if s.InitialTime.IsZero() {
		return
	}
	now := time.Now().UTC()
	total := nanoToMilli(now.Sub(s.InitialTime).Nanoseconds())
	s.Time = total
}

// GetElapsedTime get elapsed time in milliseconds
func (s *StopWatch) GetElapsedTime() float64 {
	if s.InitialTime.IsZero() {
		return 0
	}
	return nanoToMilli(time.Now().UTC().Sub(s.InitialTime).Nanoseconds())
}

// GetElapsedAndReset convenience method to get elapsed time and reset right away
func (s *StopWatch) GetElapsedAndReset() float64 {
	elapsed := s.GetElapsedTime()
	s.InitialTime = time.Time{}
	s.Time = 0
	return elapsed
}

func nanoToMilli(nanoTime int64) float64 {
	return float64(nanoTime) / 1000000
}