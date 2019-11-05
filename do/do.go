package do

import (
	"fmt"
	"sync"
	"time"
)

var (
	keys    = map[string]time.Time{}
	doneMut sync.Mutex
)

// ShouldDo 指定key一段时间内是否需要操作
func ShouldDo(d time.Duration, key string) bool {
	doneMut.Lock()
	defer doneMut.Unlock()
	if t, ok := keys[key]; !ok || t.Before(time.Now()) {
		keys[key] = time.Now().Add(d)
		return true
	}
	return false
}

// ShouldDof 指定key一段时间内是否需要操作, key 格式化
func ShouldDof(d time.Duration, format string, v ...interface{}) bool {
	return ShouldDo(d, fmt.Sprintf(format, v...))
}

// SetExpires 设置指定key过期时间
func SetExpires(t time.Time, key string) {
	doneMut.Lock()
	defer doneMut.Unlock()
	keys[key] = t
}

// SetExpiresf 设置指定key过期时间，key 格式化
func SetExpiresf(t time.Time, format string, v ...interface{}) {
	SetExpires(t, fmt.Sprintf(format, v...))
}

// Clear 清除指定key过期时间
func Clear(key string) {
	doneMut.Lock()
	defer doneMut.Unlock()
	delete(keys, key)
}

// Clearf 清除指定key过期时间，key 格式化
func Clearf(format string, v ...interface{}) {
	Clear(fmt.Sprintf(format, v...))
}

// SyncDone 非阻塞发送 done 消息
func SyncDone(done chan<- struct{}) {
	select {
	case done <- struct{}{}:
	default:
	}
}
