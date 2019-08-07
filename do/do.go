package fcoin

import (
	"fmt"
	"sync"
	"time"
)

var (
	items   map[string]time.Time
	doneMut sync.Mutex
)

// ShouldDo 指定key一段时间内是否需要操作
func ShouldDo(d time.Duration, key string) bool {
	doneMut.Lock()
	defer doneMut.Unlock()
	if t, ok := items[key]; !ok || t.Before(time.Now()) {
		items[key] = time.Now().Add(d)
		return true
	}
	return false
}

// ShouldDof 指定key一段时间内是否需要操作, key 格式化
func ShouldDof(d time.Duration, format string, v ...interface{}) bool {
	return ShouldDo(d, fmt.Sprintf(format, v...))
}