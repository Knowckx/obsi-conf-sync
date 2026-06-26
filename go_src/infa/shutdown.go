package infa

import (
	"sync"

	"github.com/cockroachdb/errors"
)

type CleanupFunc func() error

func RegisterCleanup(name string, fn CleanupFunc) error {
	return shutdownMgr.Register(name, fn)
}

func RunCleanup() error {
	return shutdownMgr.RunCleanup()
}

var shutdownMgr = shutdownManager{}

type shutdownManager struct {
	mu      sync.Mutex
	items   []cleanupItem
	running bool
	
	// 目前的设计是 进程退出清理中心 
	// 所以保留done 清理前可多次注册，整个进程只清理一次
	done    bool
}

type cleanupItem struct {
	name string
	fn   CleanupFunc
}

func (m *shutdownManager) Register(name string, fn CleanupFunc) error {
	if fn == nil {
		return errors.New("shutdown: cleanup func is nil")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.running {
		return errors.New("shutdown: cannot register cleanup while running")
	}
	if m.done {
		return errors.New("shutdown: cleanup already finished")
	}

	m.items = append(m.items, cleanupItem{
		name: name,
		fn:   fn,
	})
	return nil
}

func (m *shutdownManager) RunCleanup() error {
	m.mu.Lock()
	if m.running || m.done {
		m.mu.Unlock()
		return nil
	}

	m.running = true
	items := append([]cleanupItem(nil), m.items...)
	m.mu.Unlock()

	defer func() {
		m.mu.Lock()
		m.running = false
		m.done = true
		m.items = nil
		m.mu.Unlock()
	}()

	var errs []error

	// 反序执行，模拟 defer
	for i := len(items) - 1; i >= 0; i-- {
		if err := runCleanupItem(items[i]); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func runCleanupItem(item cleanupItem) (err error) {
	defer func() {
		panicErr := panicToError(recover())
		if panicErr == nil {
			return
		}
		err = errors.Wrapf(panicErr, "cleanup %s panic", item.name)
	}()

	err = item.fn()
	if err != nil {
		return errors.Wrapf(err, "cleanup %s", item.name)
	}
	return nil
}

// panic → error
func panicToError(panicValue any) error {
	if panicValue == nil {
		return nil
	}

	panicErr, ok := panicValue.(error)
	if ok {
		return errors.WithStack(panicErr)
	}
	return errors.Errorf("panic: %v", panicValue)
}
