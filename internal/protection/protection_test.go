package protection

import (
	"sync/atomic"
	"testing"
)

func Test_Protector(t *testing.T) {
	a := int64(4)

	swapped := atomic.CompareAndSwapInt64(&a, 6, 5)
	t.Log(swapped)
	t.Log(a)
}
