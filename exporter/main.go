package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/load", load)
	http.ListenAndServe(":3000", nil)
}

func load(w http.ResponseWriter, r *http.Request) {
	now := float64(time.Now().Unix())
	io.WriteString(w, fmt.Sprintf("%v", math.Abs(math.Cos(now/100))))
}
