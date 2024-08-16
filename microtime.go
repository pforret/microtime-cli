package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	diffTime := flag.Float64("diff", 0, "Calculate duration from the given time")
	formatCmd := flag.String("format", "", "Output format (options: cmd)")
	flag.Parse()
	now := time.Now()
	microTime := float64(now.Unix()) + (float64(now.Nanosecond()) / 1000000000)
	if *formatCmd == "cmd" {
		if *diffTime > 0 {
			fmt.Printf("set ELAPSED=%.6f\n", microTime-*diffTime)
		} else {
			fmt.Printf("set MICROTIME=%.6f\n", microTime)
		}
	} else {
		if *diffTime > 0 {
			fmt.Printf("%.6f\n", microTime-*diffTime)
		} else {
			fmt.Printf("%.6f\n", microTime)
		}
	}
}
