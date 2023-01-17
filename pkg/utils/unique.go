package utils

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

func GenerateUniqueCode(prefix string) string {
	mu.Lock()
	defer mu.Unlock()

	return fmt.Sprintf("%v-%v", prefix, time.Now().UnixNano())
}
