package main

import (
	"fmt"
	"time"
)

// UnixToAgo ...
func UnixToAgo(ctime int64) string {

	ago := (time.Now().Unix() - ctime)
	if ago < 60 {
		return fmt.Sprintf("%ds", ago)
	} else if ago < 3600 {
		return fmt.Sprintf("%dm", int64(ago/60))
	} else if ago < (86400) {
		return fmt.Sprintf("%dh", int64(ago/3600))
	} else if ago < (604800) {
		return fmt.Sprintf("%dd", int64(ago/86400))
	} else if ago < (2628000) {
		return fmt.Sprintf("%dw", int64(ago/604800))
	} else if ago < (31536000) {
		return fmt.Sprintf("%dM", int64(ago/2628000))
	}
	return fmt.Sprintf("%dy", int64(ago/31536000))
}

func main() {
	fmt.Println(UnixToAgo(1524221347))
}
