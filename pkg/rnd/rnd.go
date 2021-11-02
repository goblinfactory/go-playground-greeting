package rnd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// SleepMinMaxMs for min or max milliseconds
func SleepMinMaxMs(min int, max int) {
	ms := time.Duration(rand.Intn(max-min)+min) * time.Millisecond
	time.Sleep(ms)
}

// SleepMs sleeps for milliseconds
func SleepMs(ms int) {
	time.Sleep(time.Duration(ms * int(time.Millisecond)))
}

// RandFloatString return a random float string with n decimals
func RandFloatString(max int, numDecimals int) string {
	f := rand.Float64() * float64(max)
	nd := fmt.Sprintf("%d", numDecimals)
	return fmt.Sprintf("%."+nd+"f", f)
}

// RandFloat return what we hope is random float string with n decimals. (not guaranteed with floats, but hey, good enough for demo code!)
func RandFloat(max int, numdecimals int) float64 {
	s := RandFloatString(max, numdecimals)
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
