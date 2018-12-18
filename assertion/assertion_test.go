package assertion

import (
	"testing"
)

type Executor1 interface {
	Exec() error
}

type CTExecutor struct {
}

func (cte *CTExecutor) Exec() error {
	return nil
}

var cte = &CTExecutor{}

func BenchmarkAssertionSubInterface2(b *testing.B) {
	var e Executor1 = cte
	for i := 0; i < b.N; i++ {
		_ = e.(*CTExecutor)
	}
}

func BenchmarkAssertionInterface2(b *testing.B) {
	var e interface{} = cte
	for i := 0; i < b.N; i++ {
		_ = e.(*CTExecutor)
	}
}
