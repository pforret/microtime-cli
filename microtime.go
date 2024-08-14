package main

import (
    "flag"
    "fmt"
    "time"
)

func main() {
    diffTime := flag.Float64("diff",0,"Calculate duration from the given time")
    flag.Parse()
    now := time.Now()
    microTime := float64(now.Unix()) + (float64(now.Nanosecond()) / 1000000000)
    if(*diffTime > 0) {
         fmt.Printf("%.6f\n",microTime - *diffTime)
   } else {
        fmt.Printf("%.6f\n",microTime)
    }
}
