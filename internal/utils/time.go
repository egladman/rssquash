package utils

import (
	"time"
)

func GetCurrentTime() string {
	t := time.Now().Format(time.RFC3339)
	return t

}
