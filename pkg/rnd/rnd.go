package rnd

import (
	"math/rand"
	"time"
)

// Sleep for min or max milliseconds
func Sleep(min int, max int) {
	ms := time.Duration(rand.Intn(min+max)) * time.Millisecond
	time.Sleep(ms)
}

// SleepMs sleeps for milliseconds
func SleepMs(ms int) {
	time.Sleep(time.Duration(ms * int(time.Millisecond)))
}
