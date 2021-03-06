package core

// RuntimeFigures are meta values given by respective impl
type RuntimeFigures map[string]float64

const (
	// Iterations is the number of iterations
	Iterations = "iterations"
	// PushedData is the number of pushed data
	PushedData = "pushedData"
	// Duration is algo duration
	Duration = "duration"
	// LastDataTime is the last pushed data time
	LastDataTime = "lastDataTime"
)
