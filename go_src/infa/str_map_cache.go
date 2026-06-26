package infa

import (
	"sync"

	"github.com/cockroachdb/errors"
)

// StrCache 是一个字符串键控的并发缓存。
type StrCache[T any] struct {
	mu    sync.RWMutex
	items map[string]T
}

// NewStrCache 返回一个空的字符串键控缓存。
func NewStrCache[T any]() *StrCache[T] {
	return &StrCache[T]{
		items: map[string]T{},
	}
}

// GetKey 返回 key 对应的值；不存在时 ok=false。
func (r *StrCache[T]) GetKey(key string) (T, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	value, ok := r.items[key]
	return value, ok
}

// Get 返回 key 对应的值；不存在时返回 error。
func (r *StrCache[T]) Get(key string) (T, error) {
	value, ok := r.GetKey(key)
	if ok {
		return value, nil
	}

	var zero T
	return zero, errors.Errorf("infa.str-map-cache: key not found: %s", key)
}

// Set 写入 key 对应的值；如果发生覆盖，返回旧值和 replaced=true。
func (r *StrCache[T]) Set(key string, value T) (old T, replaced bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	old, replaced = r.items[key]
	r.items[key] = value
	return old, replaced
}

// Values 返回当前缓存中的所有值快照；返回顺序不保证。
func (r *StrCache[T]) Values() []T {
	r.mu.RLock()
	defer r.mu.RUnlock()

	values := make([]T, 0, len(r.items))
	for _, value := range r.items {
		values = append(values, value)
	}
	return values
}

// Keys 返回当前缓存中的所有键快照；返回顺序不保证。
func (r *StrCache[T]) Keys() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	keys := make([]string, 0, len(r.items))
	for key := range r.items {
		keys = append(keys, key)
	}
	return keys
}

// Clear 清空当前缓存中的所有值。
func (r *StrCache[T]) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.items = map[string]T{}
}
