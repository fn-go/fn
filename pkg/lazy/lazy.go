package lazy

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//type key int
//
//var requestKey key

var ids uint64

func Noop[T any]() Lazy[T] {
	return Of[T](func() T {
		var noop T
		return noop
	})
}

func Of[T any](f func() T) Lazy[T] {
	return OfE(func() (T, error) {
		return f(), nil
	})
}

func OfE[T any](initFn func() (T, error)) Lazy[T] {
	return &lazy[T]{
		new:         initFn,
		requestedBy: make(map[uint64]struct{}),
		id:          atomic.AddUint64(&ids, 1),
	}
}

type Lazy[T any] interface {
	SetInit(fn func() (T, error))
	Eval() (T, error)
}

type lazy[T any] struct {
	new         func() (T, error)
	once        sync.Once
	value       T
	err         error
	done        bool
	id          uint64
	requestedBy map[uint64]struct{}
}

func (l *lazy[T]) SetInit(fn func() (T, error)) {
	if l.done {
		panic(fmt.Errorf("cannot call SetInit after the lazy has been evaluated: %d", l.id))
	}
	l.new = fn
}

func (l *lazy[T]) Eval() (T, error) {
	if l == nil {
		var noop T
		return noop, nil
	}

	if _, ok := l.requestedBy[l.id]; ok {
		var noop T
		return noop, fmt.Errorf("detected a cycle in lazy: %+v", l)
	}

	l.requestedBy[l.id] = struct{}{}
	defer delete(l.requestedBy, l.id)

	l.once.Do(func() {
		if l.new != nil {
			v, err := l.new()
			l.value = v
			l.err = err
			l.new = nil // so that f can now be GC'ed
		}

		l.done = true
	})

	return l.value, l.err
}

type Evaluator[T any] interface {
	Register(name string, lz Lazy[T]) error
	Evaluate(name string) (T, error)
}

type threadSafeEvaluator[T any] struct {
	lazies map[string]Lazy[T]
	sync.Mutex
}

func (t *threadSafeEvaluator[T]) Register(name string, lz Lazy[T]) error {
	t.Lock()
	defer t.Unlock()

	if _, exists := t.lazies[name]; exists {
		return fmt.Errorf("name collision found: %s", name)
	}

	t.lazies[name] = lz

	return nil
}

func (t *threadSafeEvaluator[T]) Evaluate(name string) (T, error) {
	t.Lock()
	defer t.Unlock()

	if _, exists := t.lazies[name]; !exists {
		var noop T
		return noop, fmt.Errorf("nothing registered by name: %s", name)
	}

	return t.lazies[name].Eval()
}

func NewThreadSafeEvaluator[T any]() Evaluator[T] {
	return &threadSafeEvaluator[T]{
		lazies: make(map[string]Lazy[T]),
	}
}
