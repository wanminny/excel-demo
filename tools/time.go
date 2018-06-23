package tools

import (
	"time"
	"fmt"
)

func Sub(start time.Time,end time.Time) string {
	delta := end.Sub(start)
	tmp := fmt.Sprintf("%f",delta.Seconds())
	return tmp
}
